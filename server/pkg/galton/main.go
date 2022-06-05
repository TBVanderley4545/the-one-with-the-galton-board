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

	columnAdjuster := 0

	if len(gb.Columns)%2 == 0 && newBall.NetOffset > 0 {
		columnAdjuster = 1
	}

	column := newBall.NetOffset - gb.StartingColumnIndex - columnAdjuster

	gb.Columns[column] = append(gb.Columns[column], &newBall)
}

func (gb *GaltonBoard) AddBalls(numberToAdd int) {
	for i := 0; i < numberToAdd; i++ {
		gb.AddBall()
	}
}

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
