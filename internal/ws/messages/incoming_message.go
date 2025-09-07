package messages

/*------------------Base Incoming Message------------------*/
// All messages from clients adhere to this structure
type IncomingMessage struct {
	Action string `json:"action"`            // what the client wants to do
	GameID string `json:"game_id,omitempty"` // which game
	Data   any    `json:"data,omitempty"`    // action-specific payload
}

/* ------------------Specific Messages Incoming------------------*/

// Player wants to join a game
type JoinGamePayload struct {
	Name string `json:"name"`
}

// Player left the game
type LeaveGamePayload struct {}

// Player rolls the dice
type DiceRollPayload struct {} // no dice value - server generates it

// Player wants to move a peice
type MovePiecePayload struct {
	PieceID int `json:"piece_id"` // which piece to move (0-3)
}


