package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/tarm/serial"
)

type ControlCodes struct {
	SoftwareArm         string `yaml:"software_arm"`
	LaunchSystemsArm    string `yaml:"launch_systems_arm"`
	VPRocketsArm        string `yaml:"vp_rockets_arm"`
	SoftwareLaunch      string `yaml:"software_launch"`
	LaunchSystemsLaunch string `yaml:"launch_systems_launch"`
	VPRocketsLaunch     string `yaml:"vp_rockets_launch"`
	Abort               string `yaml:"abort"`
}

type LaunchStatus struct {
	Type string `json:"type"`
	// all counters from go from 0 -> counterMax
	SoftwareArmCounter int `json:"softwareArmCounter"`
	softwareArmActive  bool

	LaunchSystemsArmCounter int `json:"launchSystemsArmCounter"`
	launchSystemsArmActive  bool

	VPRocketsArmCounter int `json:"vpRocketsArmCounter"`
	vpRocketsArmActive  bool

	ArmCounter int `json:"armCounter"`
	armed      bool

	SoftwareLaunchCounter int `json:"softwareLaunchCounter"`
	softwareLaunchActive  bool

	LaunchSystemsLaunchCounter int `json:"launchSystemsLaunchCounter"`
	launchSystemsLaunchActive  bool

	VPRocketsLaunchCounter int `json:"vpRocketsLaunchCounter"`
	vpRocketsLaunchActive  bool

	LaunchCounter int `json:"launchCounter"`

	// countdown goes from 10 -> 0
	Countdown int `json:"countdown"`
	launched  bool
}

var (
	launchStatus LaunchStatus
	controlCodes ControlCodes
	controlMutex sync.Mutex
)

const (
	counterMax  = 100
	counterTick = 10
)

func sendLaunchStatus(hub *Hub, interval time.Duration, serialConn *serial.Port) {
	go updateLaunchCounters(serialConn)

	updatePeriod, _ := time.ParseDuration("300ms")
	tick := time.NewTicker(updatePeriod)
	for {
		<-tick.C // Block until next cycle
		// log.Println("Sending LaunchStatus")
		err := hub.sendMsg(launchStatus)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func updateLaunchCounters(serialConn *serial.Port) {
	updatePeriod, _ := time.ParseDuration("1000ms")
	tick := time.NewTicker(updatePeriod)
	for {
		<-tick.C // Block until next cycle
		if launchStatus.softwareArmActive {
			launchStatus.SoftwareArmCounter += counterTick
			log.Println("SoftwareArmCounter", launchStatus.SoftwareArmCounter)
		} else if launchStatus.ArmCounter < counterMax {
			launchStatus.SoftwareArmCounter = 0
		}

		if launchStatus.launchSystemsArmActive {
			launchStatus.LaunchSystemsArmCounter += counterTick
		} else if launchStatus.ArmCounter < counterMax {
			launchStatus.LaunchSystemsArmCounter = 0
		}

		if launchStatus.vpRocketsArmActive {
			launchStatus.VPRocketsArmCounter += counterTick
		} else if launchStatus.ArmCounter < counterMax {
			launchStatus.VPRocketsArmCounter = 0
		}

		if launchStatus.SoftwareArmCounter >= counterMax &&
			launchStatus.LaunchSystemsArmCounter >= counterMax &&
			launchStatus.VPRocketsArmCounter >= counterMax {
			launchStatus.ArmCounter += counterTick
		} else if launchStatus.Countdown > 0 {
			launchStatus.ArmCounter = 0
		}

		if launchStatus.ArmCounter < counterMax {
			launchStatus.SoftwareLaunchCounter = 0
			launchStatus.LaunchSystemsLaunchCounter = 0
			launchStatus.VPRocketsLaunchCounter = 0
			launchStatus.LaunchCounter = 0
			launchStatus.Countdown = 10
			continue
		} else if !launchStatus.armed {
			launchStatus.armed = true
			serialConn.Write([]byte{0x21})
		}

		if launchStatus.softwareLaunchActive {
			launchStatus.SoftwareLaunchCounter += counterTick
		} else if launchStatus.Countdown > 0 {
			launchStatus.SoftwareLaunchCounter = 0
		}

		if launchStatus.launchSystemsLaunchActive {
			launchStatus.LaunchSystemsLaunchCounter += counterTick
		} else if launchStatus.Countdown > 0 {
			launchStatus.LaunchSystemsLaunchCounter = 0
		}

		if launchStatus.vpRocketsLaunchActive {
			launchStatus.VPRocketsLaunchCounter += counterTick
		} else if launchStatus.LaunchCounter < counterMax {
			launchStatus.VPRocketsLaunchCounter = 0
		}

		if launchStatus.SoftwareLaunchCounter >= counterMax &&
			launchStatus.LaunchSystemsLaunchCounter >= counterMax &&
			launchStatus.VPRocketsLaunchCounter >= counterMax {

			launchStatus.LaunchCounter += counterTick
		} else if launchStatus.Countdown > 0 {
			launchStatus.LaunchCounter = 0
		}

		if launchStatus.LaunchCounter >= counterMax && launchStatus.Countdown > 0 {
			launchStatus.Countdown--
			if launchStatus.Countdown <= 0 && !launchStatus.launched {
				launchStatus.launched = true
				serialConn.Write([]byte{0x20})
			}
		} else if launchStatus.Countdown > 0 {
			launchStatus.Countdown = 10
		}
	}
}

func setControlCodes(codes ControlCodes) error {
	if err := validateControlCode(codes.SoftwareArm); err != nil {
		return err
	}
	if err := validateControlCode(codes.LaunchSystemsArm); err != nil {
		return err
	}
	if err := validateControlCode(codes.VPRocketsArm); err != nil {
		return err
	}
	if err := validateControlCode(codes.SoftwareLaunch); err != nil {
		return err
	}
	if err := validateControlCode(codes.LaunchSystemsLaunch); err != nil {
		return err
	}
	if err := validateControlCode(codes.VPRocketsLaunch); err != nil {
		return err
	}
	if err := validateControlCode(codes.Abort); err != nil {
		return err
	}

	controlCodes = codes
	return nil
}

func validateControlCode(code string) error {
	if len(code) < 3 {
		return fmt.Errorf("Control code '%s' too short, must be at least 3 characters", code)
	} else if len(code) > 25 {
		return fmt.Errorf("Control code '%s' too long, must be at most 25 characters", code)
	}
	return nil
}

func handleLaunchControl(code string, controlType string) {
	controlMutex.Lock()
	defer controlMutex.Unlock()

	if controlType == "arm" || controlType == "stopArm" {
		newState := false
		if controlType == "arm" {
			newState = true
		}

		switch code {
		case controlCodes.SoftwareArm:
			log.Println("launchStatus.softwareArmActive = ", newState)
			launchStatus.softwareArmActive = newState
		case controlCodes.LaunchSystemsArm:
			log.Println("launchStatus.launchSystemsArmActive = ", newState)
			launchStatus.launchSystemsArmActive = newState
		case controlCodes.VPRocketsArm:
			log.Println("launchStatus.vpRocketsArmActive = ", newState)
			launchStatus.vpRocketsArmActive = newState
		}
	}

	if controlType == "launch" || controlType == "stopLaunch" {
		newState := false
		if controlType == "launch" {
			newState = true
		}

		switch code {
		case controlCodes.SoftwareLaunch:
			log.Println("launchStatus.softwareLaunchActive = ", newState)
			launchStatus.softwareLaunchActive = newState
		case controlCodes.LaunchSystemsLaunch:
			log.Println("launchStatus.launchSystemsLaunchActive = ", newState)
			launchStatus.launchSystemsLaunchActive = newState
		case controlCodes.VPRocketsLaunch:
			log.Println("launchStatus.vpRocketsLaunchActive = ", newState)
			launchStatus.vpRocketsLaunchActive = newState
		}
	}

	if controlType == "abort" && code == controlCodes.Abort {
		launchStatus.SoftwareArmCounter = 0
		launchStatus.softwareArmActive = false
		launchStatus.LaunchSystemsArmCounter = 0
		launchStatus.launchSystemsArmActive = false
		launchStatus.VPRocketsArmCounter = 0
		launchStatus.vpRocketsArmActive = false
		launchStatus.ArmCounter = 0

		launchStatus.SoftwareLaunchCounter = 0
		launchStatus.softwareLaunchActive = false
		launchStatus.LaunchSystemsLaunchCounter = 0
		launchStatus.launchSystemsLaunchActive = false
		launchStatus.VPRocketsLaunchCounter = 0
		launchStatus.vpRocketsLaunchActive = false
		launchStatus.LaunchCounter = 0

		launchStatus.Countdown = 10
	}
}
