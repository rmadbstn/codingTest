package programmers

import "sort"

var islandCheckArr [][]bool
var dirs [][]int

func SolutionUninhabitedIsland(maps []string) []int { //["X591X","X1X5X","X231X", "1XXX1"]

	islandCheckArr = make([][]bool, len(maps))
	dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	stayArr := []int{}

	for i := range islandCheckArr {
		islandCheckArr[i] = make([]bool, len(maps[0]))
	}

	for i := 0; i < len(islandCheckArr); i++ {

		for j := 0; j < len(islandCheckArr[i]); j++ {

			if !islandCheckArr[i][j] && maps[i][j] != 'X' {

				stayArr = append(stayArr, bfsIsland(i, j, maps))

			}

		}

	}

	sort.Ints(stayArr)

	if len(stayArr) == 0 {
		return []int{-1}
	}

	return stayArr
}

func bfsIsland(y int, x int, maps []string) int {

	queue := [][]int{}

	queue = append(queue, []int{y, x})

	// count, err := strconv.Atoi(string(maps[y][x]))
	// if err != nil {
	// 	// 변환에 실패한 경우 에러 처리
	// 	// 예를 들어, count를 0으로 설정하거나 에러 메시지 출력
	// 	count = 0
	// }
	count := int(maps[y][x] - '0')

	islandCheckArr[y][x] = true

	for len(queue) > 0 {

		curY := queue[0][0]
		curX := queue[0][1]

		queue = queue[1:]

		for _, dir := range dirs {

			newX := curX + dir[0]
			newY := curY + dir[1]

			if newY >= 0 && newX >= 0 && newY < len(islandCheckArr) && newX < len(islandCheckArr[0]) && !islandCheckArr[newY][newX] && maps[newY][newX] != 'X' {

				queue = append(queue, []int{newY, newX})
				islandCheckArr[newY][newX] = true
				count += int(maps[newY][newX] - '0')
			}

		}

	}

	return count
}
