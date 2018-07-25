package pointcalculate

import "testing"

func Test_PointCalculate_Input_Member_NonRank_CurrentPoint_0_Amount_1000_Should_Be_0(t *testing.T) {
	Amount := 1000.00
	expected := 0
	member := Member{
		Level:        "Non-Rank",
		CurrentPoint: 0,
	}
	actual := member.PointCalculate(Amount)
	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func Test_PointCalculate_Input_Member_Sliver_CurrentPoint_20_Amount_3000_Should_Be_35(t *testing.T) {
	Amount := 3000.00
	member := Member{
		Level:        "Silver",
		CurrentPoint: 20,
	}
	expected := 50
	actual := member.PointCalculate(Amount)
	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func Test_PointCalculate_Input_Member_Regular_CurrentPoint_5_Amount_2500_Should_Be_17(t *testing.T) {
	Amount := 2500.00
	member := Member{
		Level:        "Regular",
		CurrentPoint: 5,
	}
	expected := 17
	actual := member.PointCalculate(Amount)
	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func Test_PointCalculate_Input_Member_Gold_CurrentPoint_111_Amount_2500_Should_Be_147(t *testing.T) {
	Amount := 2500.00
	member := Member{
		Level:        "Gold",
		CurrentPoint: 111,
	}
	expected := 147
	actual := member.PointCalculate(Amount)
	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
