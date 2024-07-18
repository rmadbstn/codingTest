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

func travelDFS(checkArr []bool, tickets [][]string, start string) {
	partRoute = append(partRoute, start)

	if len(partRoute) == len(tickets)+1 {
		// 모든 티켓을 사용한 경우 travelRoute 업데이트
		travelRoute = make([]string, len(partRoute))
		copy(travelRoute, partRoute)
		return
	}

	for index, ticket := range tickets {
		if checkArr[index] {
			continue
		}

		if ticket[0] == start {
			checkArr[index] = true
			travelDFS(checkArr, tickets, ticket[1])

			if len(travelRoute) > 0 {
				return // 유효한 경로를 찾은 경우 조기 종료
			}

			checkArr[index] = false
		}

		// 만약 다 순회했는데 더이상 갈 수 있는 경로가 존재하지 않아
		if index == len(tickets)-1 {
			// 백트래킹: partRoute에서 마지막 요소 제거
			partRoute = partRoute[:len(partRoute)-1]

			// ICN에서 다시 시작
			if len(travelRoute) == 0 {
				travelDFS(checkArr, tickets, "ICN")
			}
		}
	}

	// 백트래킹: partRoute에서 마지막 요소 제거
	if len(partRoute) > 0 {
		partRoute = partRoute[:len(partRoute)-1]
	}
}
