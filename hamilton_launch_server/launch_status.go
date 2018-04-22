package main

import (
	"log"
	"sync"
	"time"
)

type LaunchStatus struct {
	Type  string `json:"type"`
	mutex sync.Mutex
	// all counters from go from 0 -> 100
	SoftwareArmCounter         int `json:"softwareArmCounter"`
	LaunchSystemsArmCounter    int `json:"launchSystemsArmCounter"`
	VPRocketsArmCounter        int `json:"vpRocketsArmCounter"`
	ArmCounter                 int `json:"armCounter"`
	SoftareLaunchCounter       int `json:"softareLaunchCounter"`
	LaunchSystemsLaunchCounter int `json:"launchSystemsLaunchCounter"`
	VPRocketsLaunchCounter     int `json:"vpRocketsLaunchCounter"`
	LaunchCounter              int `json:"launchCounter"`
	// countdown goes from 10 -> 0
	Countdown int `json:"countdown"`
}

var launchStatus LaunchStatus

func sendLaunchStatus(conns *SocketConnections, interval time.Duration) {
	tick := time.NewTicker(interval)
	for {
		<-tick.C // Block until next cycle
		log.Println("Sending LaunchStatus")
		err := conns.sendMsg(launchStatus)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
