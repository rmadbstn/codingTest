package programmers

import (
	"sort"
)

func SolutionHrEvaluation(scores [][]int) int { //[[2,2],[1,4],[3,2],[3,2],[2,1]]	4      완호의 점수:scores[0]

	answer := 0
	targetA, targetB := scores[0][0], scores[0][1]
	targetScore := targetA + targetB

	// 첫 번째 점수에 대해서 내림차순,
	// 첫 번째 점수가 같으면 두 번째 점수에 대해서 오름차순으로 정렬합니다.
	sort.Slice(scores, func(i, j int) bool {
		if scores[i][0] == scores[j][0] {
			return scores[i][1] < scores[j][1] // 두 번째 점수에 대해서 오름차순
		}
		return scores[i][0] > scores[j][0] // 첫 번째 점수에 대해서 내림차순
	})

	maxB := 0

	for _, score := range scores {
		a, b := score[0], score[1]

		// 만약 완호보다 두 점수가 모두 높은 사람이 있으면 -1을 반환합니다.
		if targetA < a && targetB < b {
			return -1
		}

		// 두 번째 점수에 대해 이전까지의 최대값을 갱신합니다.
		if b >= maxB {
			maxB = b
			// 현재 점수 합이 완호의 점수 합보다 크면 answer를 증가시킵니다.
			if a+b > targetScore {
				answer++
			}
		}
	}

	return answer + 1
}
