package programmers

import "sort"

var partRoute []string
var travelRoute []string

func SolutionTravelRoute(tickets [][]string) []string {
	// 티켓들을 정렬하여 알파벳 순으로 탐색할 수 있게 함
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] < tickets[j][1]
		}
		return tickets[i][0] < tickets[j][0]
	})

	var result []string
	visited := make([]bool, len(tickets))

	// dfs 함수 정의
	var dfs func(current string, path []string) bool
	dfs = func(current string, path []string) bool {
		if len(path) == len(tickets)+1 {
			result = path
			return true
		}

		for i, ticket := range tickets {
			if !visited[i] && ticket[0] == current {
				visited[i] = true
				if dfs(ticket[1], append(path, ticket[1])) {
					return true
				}
				visited[i] = false
			}
		}
		return false
	}

	dfs("ICN", []string{"ICN"})
	return result
}
