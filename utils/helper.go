package utils

import "github.com/username/ludo-engine/internal/constants"

func GetHomeStart(color constants.Color) int {
	switch color {
	case constants.Red:
		return constants.RedHomeStart
	case constants.Blue:
		return constants.BlueHomeStart
	case constants.Green:
		return constants.GreenHomeStart
	case constants.Yellow:
		return constants.YellowHomeStart
	default:
		return 0
	}
}


func GetEntryPosition(color constants.Color) int {
	switch color {
	case constants.Red:
		return 1
	case constants.Blue:
		return 14
	case constants.Green:
		return 27
	case constants.Yellow:
		return 40
	default:
		return 0
	}
}