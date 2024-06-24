package programmers

func SolutionFunctionDevlop(progresses []int, speeds []int) []int {
	queue := []int{timeLeft(100-progresses[0], speeds[0])}
	result := []int{}

	for i := 1; i < len(progresses); i++ {
		leftProgress := 100 - progresses[i]
		leftTime := timeLeft(leftProgress, speeds[i])

		if leftTime <= queue[0] {
			// 현재 작업을 큐에 추가
			queue = append(queue, leftTime)
		} else {
			// 현재 큐의 작업들을 결과에 추가하고 큐 비우기
			result = append(result, len(queue))
			queue = []int{leftTime} // 새로운 큐로 초기화
		}
	}

	// 마지막 작업 처리
	if len(queue) > 0 {
		result = append(result, len(queue))
	}

	return result
}

func timeLeft(leftProgress int, speed int) int {
	if leftProgress%speed == 0 {
		return leftProgress / speed
	} else {
		return leftProgress/speed + 1
	}
}
