package memberpoints

const (
	Regular = 1
	Sliver  = 2
	Gold    = 3
)

func CalculatePoints(name, memberlevel string, points, amount int) int {
	if memberlevel == "Silver" {
		return SilverMember(points, amount)
	}
	if memberlevel == "Gold" {
		return GoldMember(points, amount)
	}
	if memberlevel == "Regular" {
		return RegularMember(points, amount)
	}
	return 0

}

func SilverMember(points, amount int) int {
	points = points + ((amount / 200) * Sliver)
	return points
}

func GoldMember(points, amount int) int {
	points = points + ((amount / 200) * Gold)
	return points
}

func RegularMember(points, amount int) int {
	points = points + ((amount / 200) * Regular)
	return points
}
