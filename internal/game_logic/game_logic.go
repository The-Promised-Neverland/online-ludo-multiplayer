package gamelogic

import (
	"time"

	"github.com/username/ludo-engine/internal/constants"
	"github.com/username/ludo-engine/internal/models"
	"github.com/username/ludo-engine/utils"
)

// RollDice handles a dice roll and updates the game state
func RollDice(game *models.Game, playerID string, dice int) {
	game.LastDiceValue = dice
	game.LastDicePlayer = playerID

	if dice == 6 {
		game.ConsecutiveSixes++
		if game.ConsecutiveSixes > constants.MaxConsecutiveSixes {
			game.ConsecutiveSixes = 0
			NextTurn(game)
		}
	} else {
		game.ConsecutiveSixes = 0
		NextTurn(game)
	}
}

// MovePiece updates a piece position and move history
func MovePiece(game *models.Game, playerID string, peiceID int, from int, to int, moveType constants.MoveType) {
	player := game.Players[playerID]
	piece := &player.Pieces[peiceID]
	piece.Position = to

	game.MoveHistory = append(game.MoveHistory, models.GameMoveHistory{
		PlayerID:  playerID,
		PieceID:   peiceID,
		MoveType:  moveType,
		From:      from,
		To:        to,
		DiceValue: game.LastDiceValue,
		Timestamp: time.Now(),
	})

	completed := true
	for _, p := range player.Pieces {
		if !p.IsComplete {
			completed = false
			break
		}
	}
	if completed {
		game.Winners = append(game.Winners, playerID)
		game.Status = constants.StatusFinished
	}
}

// NextTurn advances the turn to the next player
func NextTurn(game *models.Game) {
	for i, pid := range game.TurnOrder {
		if pid == game.CurrentTurn {
			game.CurrentTurn = game.TurnOrder[(i+1)%len(game.TurnOrder)]
			return
		}
	}
}


// Kill Piece handles killing a piece and granting extra turn
func KillPiece(game *models.Game, killerID string, victimID string, pieceID int) {
	victim := game.Players[victimID]
	piece := &victim.Pieces[pieceID]

	// Reset piece to home
	piece.Position = utils.GetHomeStart(victim.Color)

	// extra turn for killer
	game.ExtraTurnFromCut = true
}


// A piece can leave home if the dice roll is 6
func LeaveHome(game *models.Game, playerID string, pieceID int) bool {
	player := game.Players[playerID]
	piece := &player.Pieces[pieceID] 

	if piece.Position != 0 { // already out of house
		return false
	}

	piece.Position  = utils.GetEntryPosition(player.Color)
	return true
}

// Check win condition
func CheckWinCondition(game *models.Game, playerID string) bool {
	player := game.Players[playerID]
	for _, piece := range player.Pieces {
		if !piece.IsComplete {
			return false
		}
	}
	return true
}