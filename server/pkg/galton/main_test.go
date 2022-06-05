package galton_test

import (
//	"fmt"
	"testing"

	galton "github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/galton"
)

// func TestAddBall(t *testing.T) {
// 	currentBoard := galton.CreateBoard(4)
// 
// 	currentBoard.AddBall()
// 	currentBoard.AddBall()
// 	currentBoard.AddBall()
// 
// 	expected := 4
// 
// 	actual := 0
// 
// 	for i, c := range currentBoard.Columns {
// 		fmt.Println(i, len(c))
// 		actual += len(c)
// 	}
// 
// 	if actual != expected {
// 		t.Errorf("Expected the board to have %d balls, but it had %d",
// 			expected,
// 			actual)
// 	}
// }

func TestCreateBoard(t *testing.T) {
	currentBoard := galton.CreateBoard(9)

	expectedDecisionDepth := 9
	expectedColumns := 10
	expectedStartingColumnIndex := -5

	if expectedDecisionDepth != currentBoard.DecisionDepth {
		t.Errorf("Expected the board to have a decision depth of %d, but it had %d",
			expectedDecisionDepth,
			currentBoard.DecisionDepth)
	}

	if expectedColumns != len(currentBoard.Columns) {
		t.Errorf("Expected the board to have %d columns, but it had %d",
			expectedColumns,
			len(currentBoard.Columns))
	}

	if expectedStartingColumnIndex != currentBoard.StartingColumnIndex {
		t.Errorf("Expected the board to have a starting column index of %d, but it had a starting index of %d",
			expectedStartingColumnIndex,
			currentBoard.StartingColumnIndex)
	}
}
