package main

import (
	"log"
	"time"
)

// mock send avionics reporting for mac
func handleIncomingSerial(hub *Hub) {
	tick := time.NewTicker(time.Second)
	counter := 0.1

	for {
		<-tick.C // Block until next cycle

		log.Println("Sending Avionics")
		err := hub.sendMsg(buildAccelGyroMagnetismMsg(counter))
		if err != nil {
			log.Println(err)
		}
		err = hub.sendMsg(buildBarometerMsg(counter))
		if err != nil {
			log.Println(err)
		}
		err = hub.sendMsg(buildGpsMsg(counter))
		if err != nil {
			log.Println(err)
		}
		err = hub.sendMsg(buildOxidizerTankPressureMsg(counter))
		if err != nil {
			log.Println(err)
		}
		err = hub.sendMsg(buildCombustionChamberPressureMsg(counter))
		if err != nil {
			log.Println(err)
		}
		err = hub.sendMsg(buildFlightPhaseMsg(counter))
		if err != nil {
			log.Println(err)
		}
		err = hub.sendMsg(buildVentStatusMsg(counter))
		if err != nil {
			log.Println(err)
		}
		err = hub.sendMsg(buildLoadCellDataMsg(counter))
		if err != nil {
			log.Println(err)
		}
		err = hub.sendMsg(buildFillValveStatusMsg(counter))
		if err != nil {
			log.Println(err)
		}

		counter += 0.1
	}
}

func buildAccelGyroMagnetismMsg(counter float64) AccelGyroMagnetismMsg {
	return AccelGyroMagnetismMsg{
		Type: "accelGyroMagnetism",
		Accel: Vec3{
			X: int32(counter * 5),
			Y: int32(counter * 9),
			Z: int32(counter * 13),
		},
		Gyro: Vec3{
			X: int32(counter * 17),
			Y: int32(counter * 21),
			Z: int32(counter * 25),
		},
		Magneto: Vec3{
			X: int32(counter * 29),
			Y: int32(counter * 33),
			Z: int32(counter * 37),
		},
	}
}

func buildBarometerMsg(counter float64) BarometerMsg {
	return BarometerMsg{
		Type:        "barometer",
		Pressure:    int32(counter * 5),
		Temperature: int32(counter * 9),
	}
}

func buildGpsMsg(counter float64) GpsMsg {
	return GpsMsg{
		Type:          "gps",
		Altitude:      int32(counter * 5),
		EpochTimeMsec: int32(counter * 9),
		Latitude:      int32(counter * 13),
		Longitude:     int32(counter * 17),
	}
}

func buildOxidizerTankPressureMsg(counter float64) OxidizerTankPressureMsg {
	return OxidizerTankPressureMsg{
		Type:     "oxidizerTankPressure",
		Pressure: int32(counter * 5),
	}
}

func buildCombustionChamberPressureMsg(counter float64) CombustionChamberPressureMsg {
	return CombustionChamberPressureMsg{
		Type:     "combustionChamberPressure",
		Pressure: int32(counter * 13),
	}
}

func buildFlightPhaseMsg(counter float64) FlightPhaseMsg {
	return FlightPhaseMsg{
		Type:        "flightPhase",
		FlightPhase: int8(counter) % 6,
	}
}

func buildVentStatusMsg(counter float64) VentStatusMsg {
	return VentStatusMsg{
		Type:          "ventStatus",
		VentValveOpen: int8(counter)%2 == 0,
	}
}

func buildLoadCellDataMsg(counter float64) LoadCellDataMsg {
	return LoadCellDataMsg{
		Type:      "loadCellData",
		TotalMass: int32(counter) * 15,
	}
}

func buildFillValveStatusMsg(counter float64) FillValveStatusMsg {
	return FillValveStatusMsg{
		Type:          "fillValveStatus",
		FillValveOpen: int8(counter)%2 == 1,
	}
}
