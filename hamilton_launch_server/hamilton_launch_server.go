package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

type Weather struct {
	AirTemperature   float64 `json:"airTemperature"`
	WindSpeed        float64 `json:"windSpeed"`
	WindDirection    float64 `json:"windDirection"`
	RelativeHumidity float64 `json:"relativeHumidity"`
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

func setUpExitSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,  // terminal interrupt (ctrl-c)
		syscall.SIGQUIT, // terminal quit (ctrl-\)
		syscall.SIGTERM, // termination
	)

	fmt.Println("Listening for exit signals...")

	go func() {
		signal := <-c
		fmt.Println("Got interrupt signal:", signal)
		fmt.Println("Shutting down Hamilton Launch Server...")
		os.Exit(0)
	}()
}

func main() {
	var connections []*websocket.Conn

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		connections = append(connections, conn)
		fmt.Println("New client connected")
	})

	port := "8000"

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fmt.Println("Listening on port:", port)

	tick := time.NewTicker(time.Second)

	go func() {
		counter := 0.1
		for {
			counter += 0.1
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

	setUpExitSignals()

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
