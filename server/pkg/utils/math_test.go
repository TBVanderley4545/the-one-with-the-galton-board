package utils_test

import (
	math "github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/utils"
	"testing"
)

func TestDouble(t *testing.T) {
	var expected float64 = 256
	actual := math.Double(128)

	if actual != expected {
		t.Errorf("Expected to receive %f, but received %f.",
			expected,
			actual)
	}
}
