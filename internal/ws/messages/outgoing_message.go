package messages

import (
	"time"

	"github.com/username/ludo-engine/internal/models"
)

/*------------------Base Message Outgoing------------------*/
// All outgoing messages adhere to this structure
type GameMessage struct {
	Action    string    `json:"action"`
	GameID    string    `json:"game_id,omitempty"`
	Data      any       `json:"data,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

/* ------------------Specific Messages Outgoing------------------*/

// Full game state (authoritative snapshot)
type GameStatePayload struct {
	GameID         string                    `json:"game_id"`
	Template       string                    `json:"template"`
	Players        []*models.Player          `json:"players"`
	CurrentTurn    string                    `json:"current_turn"`
	Status         string                    `json:"status"`
	MoveHistory    []*models.GameMoveHistory `json:"move_history"`
	LastDice       int                       `json:"last_dice_value,omitempty"`
	LastDicePlayer string                    `json:"last_dice_player,omitempty"`
}

// Player joined event
type PlayerJoinedPayload struct {
	PlayerID string `json:"player_id"`
	Name     string `json:"name"`
	Color    string `json:"color"`
}

// Player left event
type PlayerLeftPayload struct {
	PlayerID string `json:"player_id"`
}

// Game over event
type GameOverPayload struct {
	Winners []string `json:"winners"` // list of player IDs in order of winning
}
