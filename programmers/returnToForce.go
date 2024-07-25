package programmers

func SolutionRTF(n int, roads [][]int, sources []int, destination int) []int {
	// 인접 리스트로 그래프 표현
	graph := make([][]int, n+1)
	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], road[1])
		graph[road[1]] = append(graph[road[1]], road[0])
	}

	// 거리 배열 초기화 및 BFS 수행
	distances := make([]int, n+1)
	for i := range distances {
		distances[i] = -1
	}
	rtf(graph, destination, distances)

	// 결과 배열 생성
	result := make([]int, len(sources))
	for i, source := range sources {
		result[i] = distances[source]
	}

	return result
}

func rtf(graph [][]int, start int, distances []int) {
	queue := []int{start}
	distances[start] = 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[cur] {
			if distances[neighbor] == -1 {
				distances[neighbor] = distances[cur] + 1
				queue = append(queue, neighbor)
			}
		}
	}
}

// func SolutionRTF(n int, roads [][]int, sources []int, destination int) []int {

// 	result := []int{}

// 	for _, source := range sources {
// 		result = append(result, rtf(source, roads, destination))
// 	}

// 	return result
// }

// func rtf(source int, roads [][]int, destination int) int {

// 	queue := [][]int{}
// 	checkArr := make([]bool, len(roads))
// 	queue = append(queue, []int{source, 0})

// 	for len(queue) > 0 {

// 		curLoc := queue[0][0]
// 		curCount := queue[0][1]
// 		queue = queue[1:]

// 		if curLoc == destination {
// 			return curCount
// 		}

// 		for index, road := range roads {

// 			if road[0] == curLoc && !checkArr[index] {
// 				queue = append(queue, []int{road[1], curCount + 1})
// 				checkArr[index] = true
// 			}

// 			if road[1] == curLoc && !checkArr[index] {
// 				queue = append(queue, []int{road[0], curCount + 1})
// 				checkArr[index] = true
// 			}

// 		}

// 	}

// 	return -1
// }
