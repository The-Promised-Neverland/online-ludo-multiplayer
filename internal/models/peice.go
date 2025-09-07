package models

type Piece struct {
	ID         int  `json:"id"`          // 0–3 for each player
	Position   int  `json:"position"`    // -1 = home, 0 = start, >0 = on board
	IsComplete bool `json:"is_complete"` // reached goal?
}