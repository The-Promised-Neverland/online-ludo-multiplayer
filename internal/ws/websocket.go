package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity; adjust in production
	},
}

// ServeWs upgrades an HTTP request to Websocket and register the player
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, playerID string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	// Register the new player session
	hub.AddSession(playerID, conn)

	// Remove session on disconnect
	defer func() {
		hub.RemoveSession(playerID)
	}()

	// Listen messages concurrently
	go hub.Listen(playerID, conn)
}
