package programmers

func SolutionBehindBigNumber(numbers []int) []int {
	n := len(numbers)
	result := make([]int, n)

	// 스택을 사용하여 해결
	stack := make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		// 현재 숫자보다 작은 숫자는 스택에서 제거
		for len(stack) > 0 && stack[len(stack)-1] <= numbers[i] {
			stack = stack[:len(stack)-1]
		}

		// 스택이 비어있으면 -1을 할당
		if len(stack) == 0 {
			result[i] = -1
		} else {
			result[i] = stack[len(stack)-1]
		}

		// 현재 숫자를 스택에 추가
		stack = append(stack, numbers[i])
	}

	return result
}
