package handler

import (
	gamelogic "github.com/username/ludo-engine/internal/game_logic"
	"github.com/username/ludo-engine/internal/ws"
)

// HandleDiceRoll processes a dice roll command and returns the result.
func handleDiceRoll(h *ws.Hub, playerID string, payload ws.DiceRollPayload) {
	h.Mutex.Lock()
	gamelogic.RollDice(h.Game, playerID, payload.DiceValue)
	h.Broadcast(h.Game)
}


