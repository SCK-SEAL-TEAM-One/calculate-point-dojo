package pointcalculate_test

import (
	"pointcalculate"
	"testing"
)

func Test_CalPoint_Input_Golf_Level_Non_Rank_And_Amont_1000_And_CurrentPoint_0_Should_Be_0(t *testing.T) {
	amont := 1000
	member := pointcalculate.Member{
		Name:         "Golf",
		Level:        "Non-Rank",
		CurrentPoint: 0,
	}
	expected := 0

	actual := member.CalPoint(amont)

	if expected != actual {
		t.Errorf("Should be %d but got %d", expected, actual)
	}
}

func Test_CalPoint_Input_Aun_Level_Regular_And_Amont_2500_And_CurrentPoint_5_Should_Be_17(t *testing.T) {
	amont := 2500
	member := pointcalculate.Member{
		Name:         "Aun",
		Level:        "Regular",
		CurrentPoint: 5,
	}
	expected := 17

	actual := member.CalPoint(amont)

	if expected != actual {
		t.Errorf("Should be %d but got %d", expected, actual)
	}
}

func Test_CalPoint_Input_Lek_Level_Silver_And_Amont_3000_And_CurrentPoint_20_Should_Be_50(t *testing.T) {
	amont := 3000
	member := pointcalculate.Member{
		Name:         "Lek",
		Level:        "Silver",
		CurrentPoint: 20,
	}
	expected := 50

	actual := member.CalPoint(amont)

	if expected != actual {
		t.Errorf("Should be %d but got %d", expected, actual)
	}
}
