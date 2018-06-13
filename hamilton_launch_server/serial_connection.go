// +build !darwin

package main

import (
	"log"
	"sync"
	"time"

	"github.com/tarm/serial"
)

var serialConn *serial.Port
var serialMutex sync.Mutex

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

	serialMutex.Lock()
	defer serialMutex.Unlock()
	serialConn = s
}

func sendSerialArmCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Arm command")
	serialConn.Write([]byte{0x21})
}

func sendSerialLaunchCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Launch command")
	serialConn.Write([]byte{0x20})
}

func sendSerialFillValveOpenCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Fill Valve Open command")
	serialConn.Write([]byte{0x22})
}

func sendSerialFillValveCloseCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Fill Valve Close command")
	serialConn.Write([]byte{0x23})
}

func sendSerialAbortCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Abort command")
	serialConn.Write([]byte{0x2F})
}
