package cutomerpoint

import "testing"

func Test_SetCustomerPoint_Input_Name_Golf_Level_NonRank_Point_0_Amount_1000_Should_Be_0(t *testing.T) {
	name := "Golf"
	level := "NonRank"
	point := 0
	amount := 1000
	expectedPoint := `{"name":"Golf","level":"NonRank","point":0,"amount":1000}`

	actualPoint := SetCustomerPoint(name, level, point, amount)

	if expectedPoint != actualPoint {
		t.Errorf("expect %v but got %v", expectedPoint, actualPoint)
	}
}
