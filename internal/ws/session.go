package ws

import "github.com/gorilla/websocket"

type PlayerSession struct {
	Conn *websocket.Conn
}
