// +build linux
// +build windows

package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tarm/serial"
)

const (
	accelGyroMagnetismHeaderByte     = 0x31 // ASCII '1'
	accelGyroMagnetismLength         = 1 + 9*4 + 1
	barometerHeaderByte              = 0x32 // ASCII '2'
	barometerLength                  = 1 + 2*4 + 1
	gpsHeaderByte                    = 0x33 // ASCII '3'
	gpsLength                        = 1 + 4*4 + 1
	oxidizerTankConditionsHeaderByte = 0x34 // ASCII '4'
	oxidizerTankConditionsLength     = 1 + 2*4 + 1
)

func sendAvionicsReporting(conns *SocketConnections, avionicsPort string, avionicsBaudrate int) {
	c := &serial.Config{Name: avionicsPort, Baud: avionicsBaudrate}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Println("Attempting to open serial port failed, retrying...")
		log.Println(err)
		tick := time.NewTicker(time.Second)
		for {
			s, err = serial.OpenPort(c)
			if err == nil {
				break
			} else {
				log.Println("Attempting to open serial port failed, retrying...")
				log.Println(err)
			}
			<-tick.C // Block until next cycle
		}
	}
	defer s.Close()
	log.Println("Serial port connection successful")

	buf := make([]byte, 128)
	for {
		n, err := s.Read(buf)
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
		case oxidizerTankConditionsHeaderByte:
			// oxidizerTankConditions
			log.Printf("oxidizerTankConditions report received")
			msg, err = buildOxidizerTankConditionsMsg(buf[:n])
		default:
			log.Printf("Unhandled Avionics case: %x", buf[:n])
			continue
		}

		if err != nil {
			log.Println("Failed to parse avionics report: %x", buf[:n])
			continue
		}

		log.Printf("Sending Avionics Report")
		err := conns.sendMsg(msg)
		if err != nil {
			log.Println(err)
			continue
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

func buildOxidizerTankConditionsMsg(buf []byte) (OxidizerTankConditionsMsg, error) {
	if len(buf) != oxidizerTankConditionsLength {
		return OxidizerTankConditionsMsg{}, fmt.Errorf(
			"oxidizerTankConditions length invalid, found %d, expected %d",
			len(buf),
			oxidizerTankConditionsLength)
	}
	return OxidizerTankConditionsMsg{
		Type:        "oxidizerTankConditions",
		Pressure:    int32(binary.BigEndian.Uint32(buf[1:5])),
		Temperature: int32(binary.BigEndian.Uint32(buf[5:9])),
	}, nil
}
