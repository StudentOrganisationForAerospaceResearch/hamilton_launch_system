package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type SocketConnections struct {
	mutex sync.Mutex
	conns []*websocket.Conn
}

func (sc *SocketConnections) sendMsg(msg interface{}) error {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	for _, conn := range sc.conns {
		err := (*conn).WriteJSON(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sc *SocketConnections) addConn(conn *websocket.Conn) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	sc.conns = append(sc.conns, conn)
}
