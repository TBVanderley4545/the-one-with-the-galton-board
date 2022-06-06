package galton

// Representation of a Galton Board
type GaltonBoard struct {
	DecisionDepth       int
	Columns             [][]*Ball
	StartingColumnIndex int
}

// Add a new ball to a Galton Board
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

// Add n balls to a Galton Board
func (gb *GaltonBoard) AddBalls(numberToAdd int) {
	for i := 0; i < numberToAdd; i++ {
		gb.AddBall()
	}
}
