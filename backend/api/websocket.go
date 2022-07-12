package api

import (
	"cafe/config"
	"github.com/gorilla/websocket"
	"net/http"
)

func inAllowedHosts(str string) bool {
	for _, a := range config.C.AllowedHosts {
		if a == str {
			return true
		}
	}
	return false
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return inAllowedHosts(origin)
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
}

func readPump(conn *websocket.Conn) {
	defer conn.Close()
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
