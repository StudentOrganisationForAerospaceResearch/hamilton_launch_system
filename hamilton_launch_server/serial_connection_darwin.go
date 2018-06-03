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

func sendSerialAbortCommand() {
	log.Println("Sending Abort command")
}
