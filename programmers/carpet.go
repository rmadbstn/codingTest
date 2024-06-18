package programmers

func SolutionCarpet(brown int, yellow int) []int {
	for i := 1; i <= yellow; i++ {
		if yellow%i == 0 {
			yellow_x := yellow / i
			yellow_y := i

			if (yellow_x+2)*(yellow_y+2)-yellow == brown {
				return []int{yellow_x + 2, yellow_y + 2}
			}
		}
	}

	return []int{}
}
