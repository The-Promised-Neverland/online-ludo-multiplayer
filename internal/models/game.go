package models

import (
	"time"

	"github.com/username/ludo-engine/internal/constants"
)

type Game struct {
	ID           string               `json:"id"`
	GameTemplate string               `json:"game_template"` // 2p, 3p, 4p
	IsPrivate    bool                 `json:"is_private"`
	Players      map[string]*Player   `json:"players"`
	MaxPlayers   int                  `json:"max_players"`
	Winners      []string             `json:"winners"`
	CurrentTurn  string               `json:"current_turn"` // player ID
	TurnOrder    []string             `json:"turn_order"`
	Status       constants.GameStatus `json:"status"`
	MoveHistory  []GameMoveHistory    `json:"move_history,omitempty"`

	// Dice & turn control
	LastDiceValue    int    `json:"last_dice_value"`
	LastDicePlayer   string `json:"last_dice_player"`
	ConsecutiveSixes int    `json:"consecutive_sixes"`
	ExtraTurnFromCut bool   `json:"extra_turn_from_cut"`

	TurnTimeout time.Duration `json:"turn_timeout"`
}
