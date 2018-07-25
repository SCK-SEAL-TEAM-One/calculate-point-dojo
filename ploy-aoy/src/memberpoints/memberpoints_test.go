package memberpoints_test

import (
	. "memberpoints"
	"testing"
)

func Test_CalculatePoints_Input_Golf_1000_Should_Be_0(t *testing.T) {
	expectPoints := 0

	actualPoints := CalculatePoints("Golf", "Non-Rank", 1000)

	if expectPoints != actualPoints {
		t.Errorf("expectedt %d but got %d", expectPoints, actualPoints)
	}
}
