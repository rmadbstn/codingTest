package programmers

func SolutionClothes(clothes [][]string) int {

	myMap := make(map[string]int)

	result := 1
	for _, cloth := range clothes {

		myMap[cloth[1]] += 1

	}

	for _, num := range myMap {
		result *= (num + 1)
	}

	return result - 1

}
