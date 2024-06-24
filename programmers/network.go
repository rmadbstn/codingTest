package programmers

func SolutionNetwork(n int, computers [][]int) int {

	checker := make([]bool, len(computers))

	count := 0

	for index := range computers {

		if !checker[index] {
			count++
			network(index, computers, checker)
		}
	}
	return count

}

func network(index int, computers [][]int, checker []bool) {
	checker[index] = true

	for idx, connect := range computers[index] {

		if connect == 1 && !checker[idx] {
			network(idx, computers, checker)
		}

	}

}
