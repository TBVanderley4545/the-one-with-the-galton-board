package galton

import (
	"math"
	"math/rand"
)

type Ball struct {
	Decisions []Decision
	NetOffset int
}

func CalcuateNetOffset(netMovement int) int {
	if netMovement == 0 {
		return 0
	}

	return int(math.Floor((math.Abs(float64(netMovement))+1)/2)) * int(int(math.Abs(float64(netMovement)))/netMovement)
}

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

	b.NetOffset = CalcuateNetOffset(NetMovement)
}
