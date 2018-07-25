package pointcalculate

type Member struct {
	Name         string
	Level        string
	CurrentPoint int
}

func (m *Member) CalPoint(amont int) int {
	memberLevel := map[string]int{"Non-Rank": 0, "Regular": 1, "Silver": 2, "Gold": 3, "Platinum": 4}
	levelBonus := memberLevel[m.Level]

	return m.CurrentPoint + ((amont / 200) * levelBonus)
}
