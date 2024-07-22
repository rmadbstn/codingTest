package programmers

func SolutionImmigration(n int, times []int) int64 {
	// 이분 탐색을 위한 변수 설정
	left, right := int64(1), int64(1e18)

	for left < right {
		mid := (left + right) / 2
		total := int64(0)

		// 모든 심사관이 주어진 시간(mid) 내에 처리할 수 있는 사람 수 계산
		for _, time := range times {
			total += mid / int64(time)
		}

		// 총 처리할 수 있는 사람 수가 n보다 작다면 시간을 늘려야 함
		if total < int64(n) {
			left = mid + 1
		} else {
			// 총 처리할 수 있는 사람 수가 n보다 크거나 같다면 시간을 줄여야 함
			right = mid
		}
	}

	// left가 최소 시간을 나타냄
	return left
}
