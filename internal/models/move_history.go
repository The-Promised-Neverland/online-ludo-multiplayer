package models

import (
	"time"

	"github.com/username/ludo-engine/internal/constants"
)

// GameMoveHistory stores one move event for replay/debug
type GameMoveHistory struct {
	PlayerID  string             `json:"player_id"`
	PieceID   int                `json:"piece_id"`
	MoveType  constants.MoveType `json:"move_type"` // e.g. "enter", "move", "cut", "finish"
	From      int                `json:"from"`
	To        int                `json:"to"`
	DiceValue int                `json:"dice_value"`
	Timestamp time.Time          `json:"timestamp"`
}
