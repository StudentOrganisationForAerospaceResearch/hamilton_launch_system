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
	WeatherUpdateInterval string       `yaml:"weather_update_interval"`
	GndStnPort          string       `yaml:"gnd_stn_port"`
	AvionicsPort          string       `yaml:"avionics_port"`
	Baudrate      int          `yaml:"baudrate"`
	Port                  int          `yaml:"port"`
	Codes                 ControlCodes `yaml:"control_codes"`
}

func init() {
	launchStatus = LaunchStatus{
		Type:                       "launchControlInfo",
		SoftwareArmCounter:         0,
		LaunchSystemsArmCounter:    0,
		VPRocketsArmCounter:        0,
		ArmCounter:                 0,
		SoftwareLaunchCounter:      0,
		LaunchSystemsLaunchCounter: 0,
		VPRocketsLaunchCounter:     0,
		LaunchCounter:              0,
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

	err = setControlCodes(config.Codes)
	if err != nil {
		log.Println("Control code error:", config.Codes)
		log.Println(err)
		os.Exit(1)
	}

	hub := newHub()
	setupSerialConnections(config.GndStnPort, config.AvionicsPort, config.Baudrate)

	// Send updates
	// go sendWeather(&hub, weatherUpdateInterval)
	go handleIncomingSerial(&hub)
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
