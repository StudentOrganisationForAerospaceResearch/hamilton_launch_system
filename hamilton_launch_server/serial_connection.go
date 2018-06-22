// +build !darwin

package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/tarm/serial"
)

var avionicsSerialConn *serial.Port
var gndStnSerialConn *serial.Port
var serialMutex sync.Mutex

func setupSerialConnections(
	gndstnPort string,
	avionicsPort string,
	baudrate int) {

	// AVIONICS
	logToFile("Starting avionics serial connection")
	avionicsConfig := &serial.Config{
		Name: avionicsPort,
		Baud: baudrate,
		ReadTimeout: time.Second,
	}

	avionicsSerial, err := serial.OpenPort(avionicsConfig)
	if err != nil {
		log.Println("Attempting to open avionics serial port failed, retrying...")
		// log.Println(err)
		tick := time.NewTicker(time.Second)
		for {
			avionicsSerial, err = serial.OpenPort(avionicsConfig)
			if err == nil {
				break
			} else {
				log.Println("Attempting to open gnd stn serial port failed, retrying...")
				// log.Println(err)
			}
			<-tick.C // Block until next cycle
		}
	}
	log.Println("Avionics Serial port connection successful")

	// GROUND STATION
	logToFile("Starting avionics serial connection")
	gndStnConfig := &serial.Config{
		Name: gndstnPort,
		Baud: baudrate,
		ReadTimeout: time.Second,
	}

	gndStnSerial, err := serial.OpenPort(gndStnConfig)
	if err != nil {
		log.Println("Attempting to open avionics serial port failed, retrying...")
		// log.Println(err)
		tick := time.NewTicker(time.Second)
		for {
			gndStnSerial, err = serial.OpenPort(gndStnConfig)
			if err == nil {
				break
			} else {
				log.Println("Attempting to open gnd stn serial port failed, retrying...")
				// log.Println(err)
			}
			<-tick.C // Block until next cycle
		}
	}
	log.Println("Avionics Serial port connection successful")


	serialMutex.Lock()
	defer serialMutex.Unlock()
	avionicsSerialConn = avionicsSerial
	gndStnSerialConn = gndStnSerial
}

func sendSerialArmCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Arm command")
	avionicsSerialConn.Write([]byte{0x21})
	gndStnSerialConn.Write([]byte{0x21})	
	logToFile("Sent Arm Command 0x21")
}

func sendSerialLaunchCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Launch command")
	avionicsSerialConn.Write([]byte{0x20})
	gndStnSerialConn.Write([]byte{0x20})	
	logToFile("Sent Launch Command 0x20")
}

func sendSerialPulseVentValveCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Pulse Vent Valve command")
	avionicsSerialConn.Write([]byte{0x24})
	gndStnSerialConn.Write([]byte{0x24})	
	logToFile("Sent Pulse Vent Valve Command 0x24")
}

func sendSerialFillValveOpenCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Fill Valve Open command")
	avionicsSerialConn.Write([]byte{0x22})
	gndStnSerialConn.Write([]byte{0x22})	
	logToFile("Sent Fill Valve Open Command 0x22")
}

func sendSerialFillValveCloseCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Fill Valve Close command")
	avionicsSerialConn.Write([]byte{0x23})
	gndStnSerialConn.Write([]byte{0x23})	
	logToFile("Sent Fill Valve Close Command 0x23")
}

func sendSerialAbortCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Abort command")
	avionicsSerialConn.Write([]byte{0x2F})
	gndStnSerialConn.Write([]byte{0x2F})	
	logToFile("Sent Abort Command 0x2F")
}

func logToFile(line string) {
	f, err := os.OpenFile("commands.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("[%v]: %s\n", time.Now().UTC(), line)); err != nil {
		log.Println(err)
		return
	}
}
