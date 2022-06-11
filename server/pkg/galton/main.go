package galton

import (
	"math"
)

var CurrentBoard GaltonBoard

// Create a Galton Board with a given grid depth
func CreateBoard(gridDepth int) GaltonBoard {
	columnCount := gridDepth + 1

	currentBoard := GaltonBoard{
		DecisionDepth:       gridDepth,
		Columns:             make([][]*Ball, columnCount),
		StartingColumnIndex: int(math.Ceil(float64(-columnCount / 2))),
	}

	return currentBoard
}
