package calculatepoint_test

import (
	. "calculatepoint"
	"testing"
)

func Test_CalculatePointFrom_Input_Amount_1000_Should_Be_Point_5(t *testing.T) {
	expectedPoint := 5
	amount := 1000

	actualPoint := CalculatePointFrom(amount)

	if expectedPoint != actualPoint {
		t.Errorf("Expected %d but got %d", expectedPoint, actualPoint)
	}
}
