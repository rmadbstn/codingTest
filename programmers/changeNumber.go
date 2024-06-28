package programmers

func SolutionChangeNumber(x int, y int, n int) int { // * 2 || * 3 || + n

	if x > y {
		return -1
	} else if x == y {
		return 0
	}
	return bfsChangeNumber(x, y, n)
}

func bfsChangeNumber(x int, y int, n int) int {

	checkArr := make([]bool, y+1)

	queue := [][]int{}
	count := 0

	queue = append(queue, []int{x, count})

	for len(queue) > 0 {

		cur := queue[0]
		queue = queue[1:]

		curNum := cur[0]
		curCount := cur[1]

		for i := 0; i < 3; i++ {

			var newNum int
			newCount := curCount + 1

			switch i {
			case 1:
				newNum = curNum * 2
			case 2:
				newNum = curNum * 3
			default:
				newNum = curNum + n

			}

			if newNum == y {
				return newCount
			} else if newNum < y && !checkArr[newNum] {
				checkArr[newNum] = true
				queue = append(queue, []int{newNum, newCount})
			}

		}

	}

	return -1

}
