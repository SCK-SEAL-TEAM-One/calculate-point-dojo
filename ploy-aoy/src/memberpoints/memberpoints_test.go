package memberpoints_test

import (
	. "memberpoints"
	"testing"
)

func Test_CalculatePoints_Input_Name_Golf_MemberLevel_Non_Rank_Points_0_Amount_1000_Should_Be_0(t *testing.T) {
	expectPoints := 0

	actualPoints := CalculatePoints("Golf", "Non-Rank", 0, 1000)

	if expectPoints != actualPoints {
		t.Errorf("expectedt %d but got %d", expectPoints, actualPoints)
	}
}

func Test_CalculatePoints_Input_Name_Lek_MemberLevel_Silver_Points_20_Amount_3000_Should_Be_50(t *testing.T) {
	expectPoints := 50

	actualPoints := CalculatePoints("Lek", "Silver", 20, 3000)

	if expectPoints != actualPoints {
		t.Errorf("expectedt %d but got %d", expectPoints, actualPoints)
	}
}

func Test_CalculatePoints_Input_Name_Aun_MemberLevel_Regular_Points_5_Amount_2500_Should_Be_17(t *testing.T) {
	expectPoints := 17

	actualPoints := CalculatePoints("Aun", "Regular", 5, 2500)

	if expectPoints != actualPoints {
		t.Errorf("expectedt %d but got %d", expectPoints, actualPoints)
	}
}
