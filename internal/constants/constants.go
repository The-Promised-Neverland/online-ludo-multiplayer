package constants

import "time"

// -------------------- Game Constants --------------------
const (
	BoardSize           = 52               // total board cells
	PiecesPerPlayer     = 4                // pieces per player
	HomeLength          = 6                // cells in the final home stretch
	MaxConsecutiveSixes = 2                // max allowed consecutive sixes
	TurnTimeout         = 30 * time.Second // example turn timeout
)

// -------------------- Player Colors --------------------
type Color string

const (
	Red    Color = "red"
	Blue   Color = "blue"
	Green  Color = "green"
	Yellow Color = "yellow"
)

// Home start positions
const (
	RedHomeStart    = 101
	BlueHomeStart   = 201
	GreenHomeStart  = 301
	YellowHomeStart = 401
)

// Home end positions
const (
	RedHomeEnd    = RedHomeStart + HomeLength - 1
	BlueHomeEnd   = BlueHomeStart + HomeLength - 1
	GreenHomeEnd  = GreenHomeStart + HomeLength - 1
	YellowHomeEnd = YellowHomeStart + HomeLength - 1
)

// -------------------- Game Status --------------------
type GameStatus string

const (
	StatusWaiting   GameStatus = "waiting"   // waiting for players
	StatusOngoing   GameStatus = "ongoing"   // game in progress
	StatusFinished  GameStatus = "finished"  // game completed
	StatusAbandoned GameStatus = "abandoned" // game ended due to player leave/timeout
)

// -------------------- Move Types --------------------
type MoveType string

const (
	MoveEnter  MoveType = "enter"  // piece comes out of home
	MoveMove   MoveType = "move"   // normal forward move
	MoveCut    MoveType = "cut"    // captured opponent piece
	MoveFinish MoveType = "finish" // reached final cell
)

// -------------------- Player Actions / Events --------------------
type PlayerAction string

const (
	ActionDiceRoll  PlayerAction = "dice_roll"  // player rolled a dice
	ActionJoinGame  PlayerAction = "join_game"  // player joined the room
	ActionLeaveGame PlayerAction = "leave_game" // player left
	ActionGameState PlayerAction = "game_state" // full game state update
	ActionMovePiece PlayerAction = "move_piece" // player moved a piece
)
