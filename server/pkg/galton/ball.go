package galton

import (
	"math"
	"math/rand"
)

// Representation of a ball on a Galton Board
type Ball struct {
	Decisions []Decision
	NetOffset int
}

// Set the net offset of a ball from the furthest left column of a Galton Board
func (b *Ball) SetNetOffset(netMovement int) {
	if netMovement == 0 {
		b.NetOffset = 0

		return
	}

	b.NetOffset = int(math.Floor((math.Abs(float64(netMovement))+1)/2)) * int(int(math.Abs(float64(netMovement)))/netMovement)
}

// Simulate dropping a ball onto a Galton Board that has a grid of n depth
func (b *Ball) Drop(decisionCount int) {
	NetMovement := 0

	for i := 0; i < decisionCount; i++ {
		if rand.Intn(2) == 0 {
			b.Decisions = append(b.Decisions, Left)
			NetMovement += int(Left)
		} else {
			b.Decisions = append(b.Decisions, Right)
			NetMovement += int(Right)
		}
	}

	b.SetNetOffset(NetMovement)
}
