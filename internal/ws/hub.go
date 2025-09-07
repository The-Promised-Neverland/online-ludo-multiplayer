package ws

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/username/ludo-engine/internal/constants"
	"github.com/username/ludo-engine/internal/models"
	"github.com/username/ludo-engine/internal/ws/messages"
)

type Hub struct {
	Game     *models.Game
	Sessions map[string]*PlayerSession // playerID → session
	Mutex    sync.Mutex
}

/* -------------------- Session Management -------------------- */

// Add a player session to hub
func (h *Hub) AddSession(playerID string, conn *websocket.Conn) {
	h.Mutex.Lock()
	h.Sessions[playerID] = &PlayerSession{Conn: conn}
	h.Mutex.Unlock()
}

// Remove a player session from hub
func (h *Hub) RemoveSession(playerID string) {
	h.Mutex.Lock()
	delete(h.Sessions, playerID)
	h.Mutex.Unlock()
}

/* -------------------- Broadcasting -------------------- */

// Send a message to all connected player
func (h *Hub) Broadcast(v any) {
	h.Mutex.Lock()
	data, err := json.Marshal(v)
	if err != nil {
		return
	}
	for _, session := range h.Sessions {
		session.Conn.WriteMessage(websocket.TextMessage, data)
	}
	h.Mutex.Unlock()
}

/* -------------------- Listener -------------------- */
// Listen for incoming messages from a player
func (h *Hub) Listen(playerID string, conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			h.RemoveSession(playerID)
			break
		}
		// parse incoming message
		var incomingMsg messages.IncomingMessage
		if err := json.Unmarshal(msg, &incomingMsg); err != nil {
			continue // ignore malformed messages
		}

		// call appropriate handler based on action
		switch incomingMsg.Action {
		case "join_game":
			var payload messages.JoinGamePayload
			if err := json.Unmarshal(msg, &payload); err == nil {
				// TODO: handle join game
				h.Broadcast(messages.GameMessage{
					Action: string(constants.ActionJoinGame),
					GameID: h.Game.ID,
					Data: messages.PlayerJoinedPayload{
						PlayerID: playerID,
						Name:     payload.Name,
					},
					Timestamp: time.Now(),
				})
			}

		case "leave_game":
			// TODO: handle leave game
			h.Broadcast(messages.GameMessage{
				Action: string(constants.ActionLeaveGame),
				GameID: h.Game.ID,
				Data: messages.PlayerLeftPayload{
					PlayerID: playerID,
				},
				Timestamp: time.Now(),
			})

		case "dice_roll":
			// TODO: handle dice roll

			var payload messages.DiceRollPayload
			if err := json.Unmarshal(msg, &payload); err == nil {
				h.Broadcast(messages.GameMessage{
					Action:    string(constants.ActionDiceRoll),
					GameID:    h.Game.ID,
					Data:      nil, // no additional data needed
					Timestamp: time.Now(),
				})
			}

		case "move_piece":
			var payload messages.MovePiecePayload
			if err := json.Unmarshal(msg, &payload); err == nil {
				h.Broadcast(messages.GameMessage{
					Action:    string(constants.ActionMovePiece),
					GameID:    h.Game.ID,
					Data:      payload,
					Timestamp: time.Now(),
				})
			}
		}
	}
}
