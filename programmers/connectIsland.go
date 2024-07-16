package programmers

import (
	"sort"
)

func SolutionConnectIsland(n int, costs [][]int) int {
	answer := 0
	isVisited := make([]bool, n)
	isBridge := make([]bool, len(costs))

	// 비용이 작은 간선을 순서로 정렬
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})

	num := 0

	// 처음에 가장 작은 간선을 무조건 넣는다. 비용이 가장 작으므로
	isVisited[costs[0][0]] = true
	isVisited[costs[0][1]] = true
	isBridge[0] = true
	answer += costs[0][2]
	num++

	// 간선의 개수가 노드의 개수-1을 만족할 때까지 순회한다.
	for num < n-1 {
		for i := 1; i < len(costs); i++ {
			start, end, cost := costs[i][0], costs[i][1], costs[i][2]
			// 다리가 건설되어 있지 않고 한쪽 노드만 방문한 경우를 찾는다.
			if !isBridge[i] &&
				((!isVisited[start] && isVisited[end]) || (isVisited[start] && !isVisited[end])) {
				num++
				answer += cost
				isVisited[start] = true
				isVisited[end] = true
				isBridge[i] = true
				break
			}
		}
	}

	return answer
}

// func SolutionConnectIsland(n int, costs [][]int) int {
// 	answer := 0
// 	isVisited := make([]bool, n)
// 	isBridge := make([]bool, len(costs))

// 	// 비용이 작은 간선을 순서로 정렬
// 	sort.Slice(costs, func(i, j int) bool {
// 		return costs[i][2] < costs[j][2]
// 	})

// 	num := 0

// 	// 처음에 가장 작은 간선을 무조건 넣는다. 비용이 가장 작으므로
// 	isVisited[costs[0][0]] = true
// 	isVisited[costs[0][1]] = true
// 	isBridge[0] = true
// 	answer += costs[0][2]
// 	num++

// 	// 간선의 개수가 노드의 개수-1을 만족할 때까지 순회한다.
// 	for num < n-1 {
// 		for i := 1; i < len(costs); i++ {
// 			start, end, cost := costs[i][0], costs[i][1], costs[i][2]
// 			// 다리가 건설되어 있지 않고 한쪽 노드만 방문한 경우를 찾는다.
// 			if !isBridge[i] &&
// 				((!isVisited[start] && isVisited[end]) || (isVisited[start] && !isVisited[end])) {
// 				num++
// 				answer += cost
// 				isVisited[start] = true
// 				isVisited[end] = true
// 				isBridge[i] = true
// 				break
// 			}
// 		}
// 	}

// 	return answer
// }
