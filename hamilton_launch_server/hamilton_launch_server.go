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

	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v2"
)

type Weather struct {
	AirTemperature   float64 `json:"airTemperature"`
	WindSpeed        float64 `json:"windSpeed"`
	WindDirection    float64 `json:"windDirection"`
	RelativeHumidity float64 `json:"relativeHumidity"`
}

type Config struct {
	UpdateInterval string `yaml:"update_interval"`
	Port           int    `yaml:"port"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

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

	log.Println("Listening for exit signals...")

	go func() {
		signal := <-c
		log.Println("Got interrupt signal:", signal)
		log.Println("Shutting down Hamilton Launch Server...")
		os.Exit(0)
	}()
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Println(err)
		log.Println("Failed to load config.yml file")
		os.Exit(1)
	}

	setUpExitSignals()

	var connections []*websocket.Conn

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		connections = append(connections, conn)
		log.Println("New client connected")
	})

	updateInterval, err := time.ParseDuration(config.UpdateInterval)
	if err != nil {
		log.Println("Failed to parse update interval")
		log.Println(err)
		os.Exit(1)
	}
	tick := time.NewTicker(updateInterval)

	go func() {
		counter := 0.1
		for {
			counter += 0.1
			log.Println("sending message...")
			for _, conn := range connections {
				conn.WriteJSON(Weather{
					AirTemperature:   counter,
					WindSpeed:        counter * 2,
					WindDirection:    counter * 3,
					RelativeHumidity: counter * 4,
				})
			}
			<-tick.C // Block until next cycle
		}
	}()

	log.Println("Listening on port:", config.Port)
	err = http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
