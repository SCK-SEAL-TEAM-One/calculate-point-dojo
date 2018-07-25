package pointcalculate

type Member struct {
	Level        string
	CurrentPoint int
}

func (m Member) PointCalculate(amount float64) int {
	return m.CurrentPoint + m.CalculatePointByLevel(int(amount))

}

func (m Member) CalculatePointByLevel(amount int) int {
	levelPointMap := map[string]int{
		"Gold":     3,
		"Regular":  1,
		"Silver":   2,
		"Non-Rank": 0,
	}

	return (amount / 200) * levelPointMap[m.Level]
}
