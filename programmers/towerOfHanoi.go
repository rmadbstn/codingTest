package programmers

// SolutionTowerOfHanoi는 주어진 원판의 수 n을 1번 기둥에서 3번 기둥으로 옮기는 최소 이동 과정을 반환합니다.
func SolutionTowerOfHanoi(n int) [][]int {
	// 이동 과정을 기록할 슬라이스 선언
	moves := [][]int{}
	// 하노이 탑 이동 과정을 재귀적으로 계산하는 함수 호출
	hanoi(n, 1, 3, 2, &moves)
	return moves
}

// hanoi 함수는 원판을 재귀적으로 이동하는 과정을 기록합니다.
// n: 이동할 원판의 수
// start: 원판이 처음 위치한 시작 기둥 번호
// end: 원판이 이동해야 할 목표 기둥 번호
// temp: 이동 과정에서 사용할 보조 기둥 번호
// moves: 이동 과정을 기록할 슬라이스 포인터
func hanoi(n int, start int, end int, temp int, moves *[][]int) {
	// 이동할 원판이 하나일 경우, 단순히 시작 기둥에서 목표 기둥으로 이동
	if n == 1 {
		// 이동 과정을 기록
		*moves = append(*moves, []int{start, end})
	} else {
		// 1단계: n-1개의 원판을 시작 기둥에서 보조 기둥으로 이동
		hanoi(n-1, start, temp, end, moves)
		// 2단계: 가장 큰 원판을 시작 기둥에서 목표 기둥으로 이동
		*moves = append(*moves, []int{start, end})
		// 3단계: 보조 기둥에 있던 n-1개의 원판을 목표 기둥으로 이동
		hanoi(n-1, temp, end, start, moves)
	}
}
