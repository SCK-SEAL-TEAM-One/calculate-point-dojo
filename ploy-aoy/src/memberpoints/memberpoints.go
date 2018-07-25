package memberpoints

func CalculatePoints(name, memberlevel string, points, amount int) int {
	if memberlevel == "Silver" {
		points = points + ((amount / 200) * 2)
		return points
	}
	if memberlevel == "Regular" {
		points = points + ((amount / 200) * 1)
		return points
	}
	if memberlevel == "Gold" {
		points = points + ((amount / 200) * 3)
		return points
	}
	return 0

}
