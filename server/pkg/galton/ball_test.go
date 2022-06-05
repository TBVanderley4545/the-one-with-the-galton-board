package galton_test

import (
	"math/rand"
	"testing"
	"time"

	galton "github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/galton"
)

func TestDrop(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	newBall := galton.Ball{}

	newBall.Drop(7)

	expectedDecisionLength := 7

	actualDecisionLength := len(newBall.Decisions)

	if expectedDecisionLength != actualDecisionLength {
		t.Errorf("Expected the ball to have a decision slice with a length of %d, but it had %d",
			expectedDecisionLength,
			actualDecisionLength)
	}
}

func TestCalcuateNetOffset(t *testing.T) {
	expected := -4
	actual := galton.CalcuateNetOffset(3)

	if expected != actual {
		t.Errorf("Expected the net offset to be %d, but it was %d",
			expected,
			actual)
	}
}
