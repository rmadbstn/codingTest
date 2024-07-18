package programmers

func SolutionFarthestNode(n int, edge [][]int) int { //[[3, 6], [4, 3], [3, 2], [1, 3], [1, 2], [2, 4], [5, 2]]    //return 3

	return fardestNodesBfs(edge)

}

func fardestNodesBfs(edge [][]int) int {
	// level별로 한번에 체크하고 queue를 비우는 식으로 하면 될듯.
	queue := []int{}
	checkMap := make(map[int]bool)
	queue = append(queue, 1)
	checkMap[1] = true

	var levelSlice []int

	for len(queue) > 0 {
		levelSlice = []int{} // levelSlice를 초기화

		curLevelSize := len(queue)
		for i := 0; i < curLevelSize; i++ {
			curNode := queue[i]

			for _, e := range edge {
				if e[0] == curNode && !checkMap[e[1]] {
					levelSlice = append(levelSlice, e[1])
					checkMap[e[1]] = true
				} else if e[1] == curNode && !checkMap[e[0]] {
					levelSlice = append(levelSlice, e[0])
					checkMap[e[0]] = true
				}
			}
		}

		// fmt.Println("새로 넣을 levelSlice갯수:" + strconv.Itoa(len(levelSlice)))

		// 큐 비우고 levelSlice채우기
		if len(levelSlice) > 0 {
			// fmt.Println("큐비우기 갯수:" + strconv.Itoa(len(queue)))
			queue = queue[:0]
			queue = append(queue, levelSlice...)
		} else {
			break
		}
	}

	return len(queue)
}
