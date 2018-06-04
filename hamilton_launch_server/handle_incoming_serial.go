// +build !darwin

package main

import (
	"encoding/binary"
	"fmt"
	"log"
)

const (
	// reference
	// https://docs.google.com/spreadsheets/d/1zXDzxP0ji3GgG9f5LK1LuefXFdcyWw12UgOviBXLQS4/edit?usp=sharing
	accelGyroMagnetismHeaderByte        = 0x31 // ASCII '1'
	accelGyroMagnetismLength            = 1 + 9*4 + 1
	barometerHeaderByte                 = 0x32 // ASCII '2'
	barometerLength                     = 1 + 2*4 + 1
	gpsHeaderByte                       = 0x33 // ASCII '3'
	gpsLength                           = 1 + 4*4 + 1
	oxidizerTankPressureHeaderByte      = 0x34 // ASCII '4'
	oxidizerTankPressureLength          = 1 + 1*4 + 1
	combustionChamberPressureHeaderByte = 0x35 // ASCII '5'
	combustionChamberPressureLength     = 1 + 1*4 + 1
	flightPhaseHeaderByte               = 0x36 // ASCII '6'
	flightPhaseLength                   = 1 + 1*1 + 1
	ventStatusHeaderByte                = 0x37 // ASCII '7'
	ventStatusLength                    = 1 + 1*1 + 1
	loadCellDataHeaderByte              = 0x40 // ASCII '7'
	loadCellDataLength                  = 1 + 1*4 + 1

	maxTotalMassKg             = 5000 // TODO
	maxOxidizerTankPressureKpa = 5660
)

func handleIncomingSerial(hub *Hub) {
	buf := make([]byte, 128)
	oxidizerTankPressure := 0
	loadCellTotalMass := 0

	for {
		n, err := serialConn.Read(buf)
		if err != nil {
			log.Println(err)
			continue
		}

		var msg interface{}
		switch buf[0] {
		case accelGyroMagnetismHeaderByte:
			// accelGyroMagnetism
			log.Printf("accelGyroMagnetism report received")
			msg, err = buildAccelGyroMagnetismMsg(buf[:n])
		case barometerHeaderByte:
			// barometer
			log.Printf("barometer report received")
			msg, err = buildBarometerMsg(buf[:n])
		case gpsHeaderByte:
			// gps
			log.Printf("gps report received")
			msg, err = buildGpsMsg(buf[:n])
		case oxidizerTankPressureHeaderByte:
			// oxidizerTankPressure
			log.Printf("oxidizerTankPressure report received")
			var tankMsg OxidizerTankPressureMsg
			tankMsg, err = buildOxidizerTankPressureMsg(buf[:n])
			oxidizerTankPressure = int(tankMsg.Pressure)
			adjustFillValve(oxidizerTankPressure, loadCellTotalMass, hub)
			msg = tankMsg
		case combustionChamberPressureHeaderByte:
			// combustionChamberPressure
			log.Printf("combustionChamberPressure report received")
			msg, err = buildCombustionChamberPressureMsg(buf[:n])
		case flightPhaseHeaderByte:
			// flightPhase
			log.Printf("flightPhase report received")
			msg, err = buildFlightPhaseMsg(buf[:n])
		case ventStatusHeaderByte:
			// flightPhase
			log.Printf("ventStatus report received")
			msg, err = buildVentStatusMsg(buf[:n])
		case loadCellDataHeaderByte:
			// flightPhase
			log.Printf("loadCellData report received")
			var massMsg LoadCellDataMsg
			massMsg, err = buildLoadCellDataMsg(buf[:n])
			loadCellTotalMass = int(massMsg.TotalMass)
			adjustFillValve(oxidizerTankPressure, loadCellTotalMass, hub)
			msg = massMsg
		default:
			log.Printf("Unhandled Avionics case: %x", buf[:n])
			continue
		}

		if err != nil {
			log.Println("Failed to parse avionics report: %x", buf[:n])
			continue
		}

		log.Printf("Sending Serial Report")
		err = hub.sendMsg(msg)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func adjustFillValve(pressure int, mass int, hub *Hub) {
	msg := FillValveStatusMsg{
		Type:          "fillValveStatus",
		FillValveOpen: false,
	}

	if pressure > maxOxidizerTankPressureKpa*0.95 ||
		mass > maxTotalMassKg {
		msg.FillValveOpen = true
		sendSerialFillValveOpenCommand()
	} else {
		sendSerialFillValveCloseCommand()
	}

	err := hub.sendMsg(msg)
	if err != nil {
		log.Println(err)
	}
}

func buildAccelGyroMagnetismMsg(buf []byte) (AccelGyroMagnetismMsg, error) {
	if len(buf) != accelGyroMagnetismLength {
		return AccelGyroMagnetismMsg{}, fmt.Errorf(
			"accelGyroMagnetism length invalid, found %d, expected %d",
			len(buf),
			accelGyroMagnetismLength)
	}
	return AccelGyroMagnetismMsg{
		Type: "accelGyroMagnetism",
		Accel: Vec3{
			X: int32(binary.BigEndian.Uint32(buf[1:5])),
			Y: int32(binary.BigEndian.Uint32(buf[5:9])),
			Z: int32(binary.BigEndian.Uint32(buf[9:13])),
		},
		Gyro: Vec3{
			X: int32(binary.BigEndian.Uint32(buf[13:17])),
			Y: int32(binary.BigEndian.Uint32(buf[17:21])),
			Z: int32(binary.BigEndian.Uint32(buf[21:25])),
		},
		Magneto: Vec3{
			X: int32(binary.BigEndian.Uint32(buf[25:29])),
			Y: int32(binary.BigEndian.Uint32(buf[29:33])),
			Z: int32(binary.BigEndian.Uint32(buf[33:37])),
		},
	}, nil
}

func buildBarometerMsg(buf []byte) (BarometerMsg, error) {
	if len(buf) != barometerLength {
		return BarometerMsg{}, fmt.Errorf(
			"barometer length invalid, found %d, expected %d",
			len(buf),
			barometerLength)
	}
	return BarometerMsg{
		Type:        "barometer",
		Pressure:    int32(binary.BigEndian.Uint32(buf[1:5])),
		Temperature: int32(binary.BigEndian.Uint32(buf[5:9])),
	}, nil
}

func buildGpsMsg(buf []byte) (GpsMsg, error) {
	if len(buf) != gpsLength {
		return GpsMsg{}, fmt.Errorf(
			"gps length invalid, found %d, expected %d",
			len(buf),
			gpsLength)
	}
	return GpsMsg{
		Type:          "gps",
		Altitude:      int32(binary.BigEndian.Uint32(buf[1:5])),
		EpochTimeMsec: int32(binary.BigEndian.Uint32(buf[5:9])),
		Latitude:      int32(binary.BigEndian.Uint32(buf[9:13])),
		Longitude:     int32(binary.BigEndian.Uint32(buf[13:17])),
	}, nil
}

func buildOxidizerTankPressureMsg(buf []byte) (OxidizerTankPressureMsg, error) {
	if len(buf) != oxidizerTankPressureLength {
		return OxidizerTankPressureMsg{}, fmt.Errorf(
			"oxidizerTankPressure length invalid, found %d, expected %d",
			len(buf),
			oxidizerTankPressureLength)
	}
	return OxidizerTankPressureMsg{
		Type:     "oxidizerTankPressure",
		Pressure: int32(binary.BigEndian.Uint32(buf[1:5])),
	}, nil
}

func buildCombustionChamberPressureMsg(buf []byte) (CombustionChamberPressureMsg, error) {
	if len(buf) != combustionChamberPressureLength {
		return CombustionChamberPressureMsg{}, fmt.Errorf(
			"combutionChamberPressure length invalid, found %d, expected %d",
			len(buf),
			combustionChamberPressureLength)
	}
	return CombustionChamberPressureMsg{
		Type:     "combustionChamberPressure",
		Pressure: int32(binary.BigEndian.Uint32(buf[1:5])),
	}, nil
}

func buildFlightPhaseMsg(buf []byte) (FlightPhaseMsg, error) {
	if len(buf) != flightPhaseLength {
		return FlightPhaseMsg{}, fmt.Errorf(
			"flightPhase length invalid, found %d, expected %d",
			len(buf),
			flightPhaseLength)
	}
	return FlightPhaseMsg{
		Type:        "flightPhase",
		FlightPhase: int8(buf[1]),
	}, nil
}

func buildVentStatusMsg(buf []byte) (VentStatusMsg, error) {
	if len(buf) != ventStatusLength {
		return VentStatusMsg{}, fmt.Errorf(
			"ventStatus length invalid, found %d, expected %d",
			len(buf),
			ventStatusLength)
	}
	return VentStatusMsg{
		Type:          "ventStatus",
		VentValveOpen: int8(buf[1]) != 0,
	}, nil
}

func buildLoadCellDataMsg(buf []byte) (LoadCellDataMsg, error) {
	if len(buf) != loadCellDataLength {
		return LoadCellDataMsg{}, fmt.Errorf(
			"loadCellData length invalid, found %d, expected %d",
			len(buf),
			loadCellDataLength)
	}
	return LoadCellDataMsg{
		Type:      "loadCellData",
		TotalMass: int32(binary.BigEndian.Uint32(buf[1:5])),
	}, nil
}
