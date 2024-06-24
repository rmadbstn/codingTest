package programmers

func SolutionProcess(priorities []int, location int) int {
	count := 0
	queue := make([]int, len(priorities))
	copy(queue, priorities)

	for {
		// 현재 문서의 중요도를 확인
		check := queue[0]
		queue = queue[1:]
		location--

		// 현재 문서보다 중요한 문서가 있는지 확인
		hasHigherPriority := false
		for _, priority := range queue {
			if priority > check {
				hasHigherPriority = true
				break
			}
		}

		if hasHigherPriority {
			// 중요도가 더 높은 문서가 있으면 현재 문서를 뒤로 보낸다
			queue = append(queue, check)
			if location < 0 {
				location = len(queue) - 1
			}
		} else {
			// 현재 문서를 출력한다
			count++
			if location < 0 {
				// 찾고자 하는 문서가 출력된 경우
				return count
			}
		}
	}
}
