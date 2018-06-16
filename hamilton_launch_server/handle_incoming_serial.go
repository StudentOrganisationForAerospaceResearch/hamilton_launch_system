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
	accelGyroMagnetismLength            = 4 + 9*4 + 1
	barometerHeaderByte                 = 0x32 // ASCII '2'
	barometerLength                     = 4 + 2*4 + 1
	gpsHeaderByte                       = 0x33 // ASCII '3'
	gpsLength                           = 4 + 4*4 + 1
	oxidizerTankPressureHeaderByte      = 0x34 // ASCII '4'
	oxidizerTankPressureLength          = 4 + 1*4 + 1
	combustionChamberPressureHeaderByte = 0x35 // ASCII '5'
	combustionChamberPressureLength     = 4 + 1*4 + 1
	flightPhaseHeaderByte               = 0x36 // ASCII '6'
	flightPhaseLength                   = 4 + 1*1 + 1
	ventStatusHeaderByte                = 0x37 // ASCII '7'
	ventStatusLength                    = 4 + 1*1 + 1
	loadCellDataHeaderByte              = 0x40 // ASCII '7'
	loadCellDataLength                  = 4 + 1*4 + 1

	maxTotalMassKg             = 5000 // TODO
	maxOxidizerTankPressureKpa = 5660
)

var (
	serialBuffer []byte
)

func handleIncomingSerial(hub *Hub) {
	buf := make([]byte, 128)

	for {
		n, err := serialConn.Read(buf)
		if err != nil {
			log.Println(err)
			continue
		}

		for i := 0; i < n; i++ {
			handleIncomingSerialByte(buf[i], hub)
		}
	}
}

func handleIncomingSerialByte(b byte, hub *Hub) {
	serialBuffer = append(serialBuffer, b)

	if serialBuffer[0] != accelGyroMagnetismHeaderByte &&
		serialBuffer[0] != barometerHeaderByte &&
		serialBuffer[0] != gpsHeaderByte &&
		serialBuffer[0] != oxidizerTankPressureHeaderByte &&
		serialBuffer[0] != combustionChamberPressureHeaderByte &&
		serialBuffer[0] != flightPhaseHeaderByte &&
		serialBuffer[0] != ventStatusHeaderByte &&
		serialBuffer[0] != loadCellDataHeaderByte {
		serialBuffer = []byte{}
		return
	} else if len(serialBuffer) == 2 {
		if serialBuffer[0] != serialBuffer[1] {
			serialBuffer = []byte{b}
			return
		}
	} else if len(serialBuffer) == 3 {
		if serialBuffer[1] != serialBuffer[2] {
			serialBuffer = []byte{b}
			return
		}
	} else if len(serialBuffer) == 4 {
		if serialBuffer[2] != serialBuffer[3] {
			serialBuffer = []byte{b}
			return
		}
	} else {
		var msg interface{}
		var err error

		// we got good header bytes
		if serialBuffer[0] == accelGyroMagnetismHeaderByte &&
			len(serialBuffer) == accelGyroMagnetismLength {
			msg, err = buildAccelGyroMagnetismMsg(serialBuffer)
		} else if serialBuffer[0] == barometerHeaderByte &&
			len(serialBuffer) == barometerLength {
			msg, err = buildBarometerMsg(serialBuffer)
		} else if serialBuffer[0] == gpsHeaderByte &&
			len(serialBuffer) == gpsLength {
			msg, err = buildGpsMsg(serialBuffer)
		} else if serialBuffer[0] == oxidizerTankPressureHeaderByte &&
			len(serialBuffer) == oxidizerTankPressureLength {
			msg, err = buildOxidizerTankPressureMsg(serialBuffer)
		} else if serialBuffer[0] == combustionChamberPressureHeaderByte &&
			len(serialBuffer) == combustionChamberPressureLength {
			msg, err = buildCombustionChamberPressureMsg(serialBuffer)
		} else if serialBuffer[0] == flightPhaseHeaderByte &&
			len(serialBuffer) == flightPhaseLength {
			msg, err = buildFlightPhaseMsg(serialBuffer)
		} else if serialBuffer[0] == ventStatusHeaderByte &&
			len(serialBuffer) == ventStatusLength {
			msg, err = buildVentStatusMsg(serialBuffer)
		} else if serialBuffer[0] == loadCellDataHeaderByte &&
			len(serialBuffer) == loadCellDataLength {
			msg, err = buildLoadCellDataMsg(serialBuffer)
		} else {
			// still reading message
			return
		}

		if serialBuffer[len(serialBuffer)-1] != 0 {
			//something is wrong
			serialBuffer = []byte{}
			return
		}

		serialBuffer = []byte{}

		if err != nil {
			log.Println("Failed to parse avionics report: %x", serialBuffer)
			return
		}

		log.Printf("Sending Serial Report")
		err = hub.sendMsg(msg)
		if err != nil {
			log.Println(err)
			return
		}

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
			X: int32(binary.BigEndian.Uint32(buf[4 : 4+4])),
			Y: int32(binary.BigEndian.Uint32(buf[8 : 8+4])),
			Z: int32(binary.BigEndian.Uint32(buf[12 : 12+4])),
		},
		Gyro: Vec3{
			X: int32(binary.BigEndian.Uint32(buf[16 : 16+4])),
			Y: int32(binary.BigEndian.Uint32(buf[20 : 20+4])),
			Z: int32(binary.BigEndian.Uint32(buf[24 : 24+4])),
		},
		Magneto: Vec3{
			X: int32(binary.BigEndian.Uint32(buf[28 : 28+4])),
			Y: int32(binary.BigEndian.Uint32(buf[32 : 32+4])),
			Z: int32(binary.BigEndian.Uint32(buf[36 : 36+4])),
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
		Pressure:    int32(binary.BigEndian.Uint32(buf[4 : 4+4])),
		Temperature: int32(binary.BigEndian.Uint32(buf[8 : 8+4])),
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
		Altitude:      int32(binary.BigEndian.Uint32(buf[4 : 4+4])),
		EpochTimeMsec: int32(binary.BigEndian.Uint32(buf[8 : 8+4])),
		Latitude:      int32(binary.BigEndian.Uint32(buf[12 : 12+4])),
		Longitude:     int32(binary.BigEndian.Uint32(buf[16 : 16+4])),
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
		Pressure: int32(binary.BigEndian.Uint32(buf[4 : 4+4])),
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
		Pressure: int32(binary.BigEndian.Uint32(buf[4 : 4+4])),
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
		FlightPhase: int8(buf[4]),
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
		VentValveOpen: int8(buf[4]) != 0,
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
		TotalMass: float64(binary.BigEndian.Int32(buf[4 : 4+4])) / 1000,
	}, nil
}
