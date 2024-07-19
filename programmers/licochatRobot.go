package programmers

func SolutionLicochatRobot(board []string) int { //["...D..R", ".D.G...", "....D.D", "D....D.", "..D...."]	,  result : 7

	// fmt.Println(string(board[0][1]))
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	checkArr := make([][]bool, len(board))

	result := -1

	for i := 0; i < len(checkArr); i++ {

		checkArr[i] = make([]bool, len(board[i]))

	}

	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[i]); j++ {

			if string(board[i][j]) == "R" {
				result = licoBFS(board, dirs, []int{i, j}, checkArr)
			}

		}

	}

	return result
}

func licoBFS(board []string, dirs [][]int, start []int, checkArr [][]bool) int {

	queue := [][]int{} // { y, x , count}

	initial := append(start, 0)

	queue = append(queue, initial)
	checkArr[start[0]][start[1]] = true
	for len(queue) > 0 {

		curX := queue[0][1]
		curY := queue[0][0]

		curCount := queue[0][2]

		queue = queue[1:]

		for _, dir := range dirs {

			newY := curY
			newX := curX

			for newY+dir[0] >= 0 && newY+dir[0] < len(board) && newX+dir[1] >= 0 && newX+dir[1] < len(board[0]) && string(board[newY+dir[0]][newX+dir[1]]) != "D" {

				newY += dir[0]
				newX += dir[1]

			}

			if string(board[newY][newX]) == "G" {
				return curCount + 1
			}

			if !checkArr[newY][newX] {
				checkArr[newY][newX] = true
				queue = append(queue, []int{newY, newX, curCount + 1})
			}

		}

	}

	return -1
}

// [
// 	"...D..R",
// 	".D.G...",
// 	"....D.D",
// 	"D....D.",
// 	"..D...."],
// 	  result : 7
