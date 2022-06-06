package galton_test

import (
	"testing"

	galton "github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/galton"
)

func TestAddBall(t *testing.T) {
	currentBoard := galton.CreateBoard(4)

	currentBoard.AddBall()
	currentBoard.AddBall()
	currentBoard.AddBall()
	currentBoard.AddBall()

	expected := 4

	actual := 0

	for _, c := range currentBoard.Columns {
		actual += len(c)
	}

	if actual != expected {
		t.Errorf("Expected the board to have %d balls, but it had %d",
			expected,
			actual)
	}
}

func TestAddBalls(t *testing.T) {
	currentBoard := galton.CreateBoard(5)
	currentBoard.AddBalls(1000)

	expected := 1000
	actual := 0

	for _, c := range currentBoard.Columns {
		actual += len(c)
	}

	if actual != expected {
		t.Errorf("Expected the board to have %d balls, but it had %d",
			expected,
			actual)
	}
}
