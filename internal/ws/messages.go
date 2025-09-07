package ws

import "github.com/username/ludo-engine/internal/constants"

// Incoming message from player
type IncomingMessage struct {
	Action  constants.PlayerAction `json:"action"`
	Payload any                    `json:"payload"`
}

// Incoming payloads
type DiceRollPayload struct {
	DiceValue int `json:"dice_value"`
}

type MovePiecePayload struct {
	PieceID int             `json:"piece_id"`
	Move    constants.MoveType `json:"move_type"`
	From    int                `json:"from"`
	To      int                `json:"to"`
}
