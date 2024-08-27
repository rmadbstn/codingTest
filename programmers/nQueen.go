package programmers

func SolutionNQueen(n int) int { //  input:4      output:2

	// 퀸을 배치하는 방법의 수를 저장할 변수
	count := 0

	// 행마다 퀸이 있는 열의 위치를 저장할 배열
	cols := make([]int, n)

	// 체스판에 퀸을 배치할 수 있는지 확인하는 함수

	canPlace := func(row, col int) bool {
		for i := 0; i < row; i++ {
			if cols[i] == col || // 같은 열에 퀸이 있는지 확인
				cols[i]-i == col-row || // 같은 \ 대각선에 퀸이 있는지 확인
				cols[i]+i == col+row { // 같은 / 대각선에 퀸이 있는지 확인
				return false
			}
		}
		return true
	}

	// 백트래킹을 이용해 퀸을 배치하는 함수
	var placeQueens func(row int)
	placeQueens = func(row int) {
		if row == n { // 모든 행에 퀸을 배치한 경우
			count++
			return
		}
		for col := 0; col < n; col++ {
			if canPlace(row, col) {
				cols[row] = col      // 퀸을 배치
				placeQueens(row + 1) // 다음 행으로 이동
			}
		}
	}

	// 백트래킹 시작
	placeQueens(0)

	// 방법의 수 반환
	return count
}
