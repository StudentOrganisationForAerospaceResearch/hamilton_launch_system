package main

import (
	"log"
)

func setupSerialConnection(avionicsPort string, avionicsBaudrate int) {
	// do nothing for mac
}

func sendSerialArmCommand() {
	log.Println("Sending Arm command")
}

func sendSerialLaunchCommand() {
	log.Println("Sending Launch command")
}

func sendSerialFillValveOpenCommand() {
	log.Println("Sending Fill Valve Open command")
}

func sendSerialFillValveCloseCommand() {
	log.Println("Sending Fill Valve Close command")
}
func sendSerialAbortCommand() {
	log.Println("Sending Abort command")
}
