package programmers

func SolutionBehindBigNumber(numbers []int) []int {

	result := make([]int, len(numbers))

	for index, number := range numbers {

		for j := index; j < len(numbers); j++ {

			if numbers[j] > number {

				result[index] = numbers[j]
				break
			}

			if j == len(numbers)-1 {
				result[index] = -1
			}
		}
	}

	return result
}
