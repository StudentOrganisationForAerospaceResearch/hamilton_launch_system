// +build !darwin

package main

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

var serialConn *serial.Port

func setupSerialConnection(avionicsPort string, avionicsBaudrate int) {
	c := &serial.Config{Name: avionicsPort, Baud: avionicsBaudrate}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Println("Attempting to open serial port failed, retrying...")
		// log.Println(err)
		tick := time.NewTicker(time.Second)
		for {
			s, err = serial.OpenPort(c)
			if err == nil {
				break
			} else {
				log.Println("Attempting to open serial port failed, retrying...")
				// log.Println(err)
			}
			<-tick.C // Block until next cycle
		}
	}
	log.Println("Serial port connection successful")
	serialConn = s
}

func sendSerialArmCommand() {
	log.Println("Sending Arm command")
	serialConn.Write([]byte{0x21})
}

func sendSerialLaunchCommand() {
	log.Println("Sending Launch command")
	serialConn.Write([]byte{0x20})
}
