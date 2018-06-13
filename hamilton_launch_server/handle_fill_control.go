package main

import (
	"log"
)

func handleFillControl(code string, cmd string, hub *Hub) {
	if code != controlCodes.FillControl {
		return
	}

	msg := FillValveStatusMsg{
		Type:          "fillValveStatus",
		FillValveOpen: false,
	}

	if cmd == "openFillValve" {
		msg.FillValveOpen = true
		sendSerialFillValveOpenCommand()
	} else if cmd == "closeFillValve" {
		msg.FillValveOpen = false
		sendSerialFillValveCloseCommand()
	}

	err := hub.sendMsg(msg)
	if err != nil {
		log.Println(err)
	}
}
