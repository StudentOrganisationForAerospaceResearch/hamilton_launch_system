package main

import (
	"log"
	"time"
)

type Weather struct {
	Type             string  `json:"type"`
	AirTemperature   float64 `json:"airTemperature"`
	WindSpeed        float64 `json:"windSpeed"`
	WindDirection    float64 `json:"windDirection"`
	RelativeHumidity float64 `json:"relativeHumidity"`
}

var counter = 0.1

func getWeather() (Weather, error) {
	counter += 0.1
	return Weather{
		Type:             "weather",
		AirTemperature:   counter,
		WindSpeed:        counter * 2,
		WindDirection:    counter * 3,
		RelativeHumidity: counter * 4,
	}, nil
}

func sendWeather(hub *Hub, interval time.Duration) {
	tick := time.NewTicker(interval)
	for {
		<-tick.C // Block until next cycle
		weather, err := getWeather()
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Sending Weather")
		err = hub.sendMsg(weather)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
