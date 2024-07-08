package programmers

func SolutionPartSeq(sequence []int, k int) []int {
	n := len(sequence)
	start, end := 0, 0
	currentSum := 0
	minLength := n + 1
	result := []int{-1, -1}

	for end < n {
		currentSum += sequence[end]

		for currentSum > k && start <= end {
			currentSum -= sequence[start]
			start++
		}

		if currentSum == k {
			if (end - start + 1) < minLength {
				minLength = end - start + 1
				result = []int{start, end}
			}
		}
		end++
	}

	return result
}
