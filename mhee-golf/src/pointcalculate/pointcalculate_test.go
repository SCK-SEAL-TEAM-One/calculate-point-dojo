package pointcalculate

import "testing"

func Test_PointCalculate_Input_Member_NonRank_CurrentPoint_0_Amount_1000_Should_Be_0(t *testing.T) {
	member := "Non-Rank"
	currentPoint := 0
	Amount := 1000.00
	expected := 0
	actual := PointCalculate(member, currentPoint, Amount)
	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
