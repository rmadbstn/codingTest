package programmers

func SolutionCreateBigNum(number string, k int) string { //("4177252841", 4)
	// runes := []rune(number)

	stack := []rune{}

	lastCount := k

	for _, target := range number {

		for lastCount > 0 && len(stack) > 0 && stack[len(stack)-1] < target {
			stack = stack[:len(stack)-1]
			lastCount--
		}

		stack = append(stack, target)

	}

	stack = stack[:len(number)-k]

	return string(stack)
}
