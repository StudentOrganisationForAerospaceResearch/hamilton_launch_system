package main

import (
	"log"
	"os"
	"time"
)

type UpdateIntervals struct {
	Weather string `yaml:"weather"`
}

// crashes on failure
func (upInt UpdateIntervals) weatherInterval() time.Duration {
	updateInterval, err := time.ParseDuration(upInt.Weather)
	if err != nil {
		log.Println("Failed to parse update interval")
		log.Println(err)
		os.Exit(1)
	}
	return updateInterval
}
