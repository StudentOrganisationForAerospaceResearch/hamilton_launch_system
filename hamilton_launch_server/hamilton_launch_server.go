package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	WeatherUpdateInterval string `yaml:"weather_update_interval"`
	AvionicsPort          string `yaml:"avionics_port"`
	AvionicsBaudrate      int    `yaml:"avionics_baudrate"`
	Port                  int    `yaml:"port"`
}

func init() {
	launchStatus = LaunchStatus{
		Type:                       "launchControlInfo",
		SoftwareArmCounter:         20,
		LaunchSystemsArmCounter:    20,
		VPRocketsArmCounter:        20,
		ArmCounter:                 20,
		SoftareLaunchCounter:       20,
		LaunchSystemsLaunchCounter: 20,
		VPRocketsLaunchCounter:     20,
		LaunchCounter:              20,
		Countdown:                  10,
	}
}

func loadConfig() (Config, error) {
	configFilename := "config.yml"
	configFile, err := ioutil.ReadFile(filepath.Join(filepath.Dir(os.Args[0]), configFilename))
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, err
	}
	log.Println("Loaded configuration:", config)
	return config, nil
}

func setUpExitSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,  // terminal interrupt (ctrl-c)
		syscall.SIGQUIT, // terminal quit (ctrl-\)
		syscall.SIGTERM, // termination
	)

	log.Println("Listening for exit signals, hit [CTRL+C] to quit")

	go func() {
		signal := <-c
		log.Println("Got interrupt signal:", signal)
		log.Println("Shutting down Hamilton Launch Server...")
		os.Exit(0)
	}()
}

func main() {
	// Setup configuration
	config, err := loadConfig()
	if err != nil {
		log.Println(err)
		log.Println("Failed to load config.yml file")
		os.Exit(1)
	}

	weatherUpdateInterval, err := time.ParseDuration(config.WeatherUpdateInterval)
	if err != nil {
		log.Println("Failed to parse update interval")
		log.Println(err)
		os.Exit(1)
	}

	hub := newHub()

	// Send updates
	go sendWeather(&hub, weatherUpdateInterval)
	go sendAvionicsReporting(&hub, config.AvionicsPort, config.AvionicsBaudrate)
	go sendFillingInfo(&hub, weatherUpdateInterval)  // use weather interval for now
	go sendLaunchStatus(&hub, weatherUpdateInterval) // use weather interval for now
	go hub.run()

	// Capture (keyboard) interrupt signals for exit
	setUpExitSignals()

	// Serve
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(&hub, w, r)
	})

	log.Println("Listening on port:", config.Port)
	err = http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
