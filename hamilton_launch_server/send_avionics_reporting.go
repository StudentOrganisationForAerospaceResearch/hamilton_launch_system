package main

import (
	"github.com/gorilla/websocket"
	"github.com/tarm/serial"
)

type Vec3 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type IMUData struct {
	Accel   Vec3 `json:"accel"`
	Gyro    Vec3 `json:"gyro"`
	Magneto Vec3 `json:"magneto"`
}

func sendAvionicsReporting(conns []*websocket.Conn, avionicsPort string, avionicsBaudrate int) {
	c := &serial.Config{Name: avionicsPort, Baud: avionicsBaudrate}
	serial.OpenPort(c)
	// s, err := serial.OpenPort(c)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// n, err := s.Write([]byte("test"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// buf := make([]byte, 128)
	// n, err = s.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%q", buf[:n])
	for {
	}
}
