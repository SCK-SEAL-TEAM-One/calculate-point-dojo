package memberpoints

func CalculatePoints(name, memberlevel string, points, amount int) int {
	if memberlevel == "Silver" {
		points = points + ((amount / 200) * 2)
		return points
	}
	return 0

}
