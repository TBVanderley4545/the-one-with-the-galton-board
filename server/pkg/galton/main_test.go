package galton_test

import (
	"testing"

	galton "github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/galton"
)

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
