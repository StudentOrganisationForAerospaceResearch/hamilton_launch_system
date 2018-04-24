package main

import (
	"log"
	"time"
)

type FillingInfo struct {
	Type          string  `json:"type"`
	TotalMass     float64 `json:"totalMass"`
	VentValveOpen bool    `json:"ventValveOpen"`
	FillValveOpen bool    `json:"fillValveOpen"`
}

var fillCounter = 1.0

func getFillingInfo() (FillingInfo, error) {
	fillCounter += 1.0
	return FillingInfo{
		Type:          "fillingInfo",
		TotalMass:     (fillCounter / 10) + 80,
		VentValveOpen: (int(fillCounter) % 2) == 0,
		FillValveOpen: (int(fillCounter) % 5) != 0,
	}, nil
}

func sendFillingInfo(hub *Hub, interval time.Duration) {
	tick := time.NewTicker(interval)
	for {
		<-tick.C // Block until next cycle
		fillingInfo, err := getFillingInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		// log.Println("Sending FillingInfo")
		err = hub.sendMsg(fillingInfo)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
