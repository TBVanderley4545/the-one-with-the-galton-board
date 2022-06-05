package galton

import (
	"math"
)

type GaltonBoard struct {
	DecisionDepth       int
	Columns             [][]*Ball
	StartingColumnIndex int
}

func (gb *GaltonBoard) AddBall() {
	newBall := Ball{}
	newBall.Drop(gb.DecisionDepth)

	column := newBall.NetOffset - gb.StartingColumnIndex

	gb.Columns[column] = append(gb.Columns[column], &newBall)
}

// Create a Galton Board with a given grid depth
func CreateBoard(gridDepth int) GaltonBoard {
	columnCount := gridDepth + 1

	currentBoard := GaltonBoard{
		DecisionDepth:       gridDepth,
		Columns:             make([][]*Ball, columnCount),
		StartingColumnIndex: int(math.Ceil(float64(columnCount/2))) * -1,
	}

	return currentBoard
}
