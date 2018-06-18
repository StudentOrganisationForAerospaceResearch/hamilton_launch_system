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

var serialConn *serial.Port
var serialMutex sync.Mutex

func setupSerialConnection(avionicsPort string, avionicsBaudrate int) {
	logToFile("Starting serial connection")
	c := &serial.Config{Name: avionicsPort, Baud: avionicsBaudrate, ReadTimeout: time.Second}
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
	logToFile("Sent Arm Command 0x21")
}

func sendSerialLaunchCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Launch command")
	serialConn.Write([]byte{0x20})
	logToFile("Sent Launch Command 0x20")
}

func sendSerialPulseVentValveCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Pulse Vent Valve command")
	serialConn.Write([]byte{0x24})
	logToFile("Sent Pulse Vent Valve Command 0x24")
}

func sendSerialFillValveOpenCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Fill Valve Open command")
	serialConn.Write([]byte{0x22})
	logToFile("Sent Fill Valve Open Command 0x22")
}

func sendSerialFillValveCloseCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Fill Valve Close command")
	serialConn.Write([]byte{0x23})
	logToFile("Sent Fill Valve Close Command 0x23")
}

func sendSerialAbortCommand() {
	serialMutex.Lock()
	defer serialMutex.Unlock()
	log.Println("Sending Abort command")
	serialConn.Write([]byte{0x2F})
	logToFile("Sent Abort Command 0x2F")
}

func logToFile(line string) {
	f, err := os.OpenFile("commands.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
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
