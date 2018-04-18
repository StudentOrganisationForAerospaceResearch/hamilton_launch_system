package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tarm/serial"
)

type Vec3 struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type AccelGyroMagnetismMsg struct {
	Type    string `json:"type"`
	Accel   Vec3   `json:"accel"`
	Gyro    Vec3   `json:"gyro"`
	Magneto Vec3   `json:"magneto"`
}

type BarometerMsg struct {
	Type        string `json:"type"`
	Pressure    int    `json:"pressure"`
	Temperature int    `json:"temperature"`
}

type GpsMsg struct {
	Type          string `json:"type"`
	Altitude      int    `json:"altitude"`
	EpochTimeMsec int    `json:"epochTimeMsec"`
	Latitude      int    `json:"latitude"`
	Longitude     int    `json:"longitude"`
}
type OxidizerTankConditionsMsg struct {
	Type        string `json:"type"`
	Pressure    int    `json:"pressure"`
	Temperature int    `json:"temperature"`
}

const (
	accelGyroMagnetismHeaderByte     = 0x31 // ASCII '1'
	barometerHeaderByte              = 0x32 // ASCII '2'
	gpsHeaderByte                    = 0x33 // ASCII '3'
	oxidizerTankConditionsHeaderByte = 0x34 // ASCII '4'
)

func sendAvionicsReporting(conns *[]*websocket.Conn, avionicsPort string, avionicsBaudrate int) {
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
			// externalPressure
			log.Printf("externalPressure report received")
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
		for _, conn := range *conns {
			conn.WriteJSON(msg)
		}
	}
}

func buildAccelGyroMagnetismMsg(buf []byte) (AccelGyroMagnetismMsg, error) {
	return AccelGyroMagnetismMsg{
		Type:    "accelGyroMagnetism",
		Accel:   Vec3{5, 5, 5},
		Gyro:    Vec3{5, 5, 5},
		Magneto: Vec3{5, 5, 5},
	}, nil
}

func buildBarometerMsg(buf []byte) (BarometerMsg, error) {
	return BarometerMsg{
		Type:        "barometer",
		Pressure:    5,
		Temperature: 5,
	}, nil
}

func buildGpsMsg(buf []byte) (GpsMsg, error) {
	return GpsMsg{
		Type:          "gps",
		Altitude:      5,
		EpochTimeMsec: 5,
		Latitude:      5,
		Longitude:     5,
	}, nil
}

func buildOxidizerTankConditionsMsg(buf []byte) (OxidizerTankConditionsMsg, error) {
	return OxidizerTankConditionsMsg{
		Type:        "oxidizerTankConditions",
		Pressure:    5,
		Temperature: 5,
	}, nil
}
