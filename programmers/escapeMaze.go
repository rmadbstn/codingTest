package programmers

func SolutionEscapeMaze(maps []string) int { //["SOOOL","XXXXO","OOOOO","OXXXX","OOOOE"]

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < len(maps); i++ {

		for j := 0; j < len(maps[i]); j++ {

			if string(maps[i][j]) == "L" {

				// fmt.Println("들어가니?")
				return escapeMaze(maps, dirs, []int{i, j})

			}

		}

	}

	return -1

}

func escapeMaze(maps []string, dirs [][]int, lever []int) int {

	distance_start := 0
	distance_finish := 0

	checkArr := make([][]bool, len(maps))
	for i := 0; i < len(checkArr); i++ {

		checkArr[i] = make([]bool, len(maps[i]))

	}

	queue := [][]int{}

	initial := append(lever, 0)

	queue = append(queue, initial)

	checkArr[lever[0]][lever[1]] = true

	for len(queue) > 0 {

		cur := queue[0]

		curY := cur[0]
		curX := cur[1]

		if string(maps[curY][curX]) == "E" && distance_finish == 0 {
			distance_finish = cur[2]
		}

		if string(maps[curY][curX]) == "S" && distance_start == 0 {
			distance_start = cur[2]
		}

		if distance_start != 0 && distance_finish != 0 {
			return distance_start + distance_finish
		}

		queue = queue[1:]

		for _, dir := range dirs {

			newY := curY + dir[0]
			newX := curX + dir[1]

			if newY >= 0 && newX >= 0 && newY < len(maps) && newX < len(maps[0]) && string(maps[newY][newX]) != "X" && !checkArr[newY][newX] {
				queue = append(queue, []int{newY, newX, cur[2] + 1})
				// fmt.Println("여기도?:", newX, newY)
				checkArr[newY][newX] = true
			}

		}

	}

	return -1
}
