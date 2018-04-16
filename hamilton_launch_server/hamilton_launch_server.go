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

type Config struct {
	UpdateIntervals UpdateIntervals `yaml:"update_intervals"`
	Port            int             `yaml:"port"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// frontend is served seperately
			return true
		},
	}
	connections []*websocket.Conn
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

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	connections = append(connections, conn)
	log.Println("New client connected")
}

func sendUpdates(updateIntervals UpdateIntervals) {
	tick := time.NewTicker(updateIntervals.weatherInterval())
	go func() {
		for {
			log.Println("Sending Weather")
			for _, conn := range connections {
				weather, err := getWeather()
				if err != nil {
					log.Println(err)
				}
				conn.WriteJSON(weather)
			}
			<-tick.C // Block until next cycle
		}
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

	http.HandleFunc("/ws", serveWs)

	sendUpdates(config.UpdateIntervals)

	log.Println("Listening on port:", config.Port)
	err = http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
