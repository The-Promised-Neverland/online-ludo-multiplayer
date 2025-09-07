package models

import "github.com/username/ludo-engine/internal/constants"

// Player
type Player struct {
	ID     string                           `json:"id"`    // unique player identifier
	Name   string                           `json:"name"`  // optional display name
	Color  constants.Color                  `json:"color"` // player’s color
	Pieces [constants.PiecesPerPlayer]Piece `json:"pieces"`
}
