package programmers

import (
	"sort"
	"strconv"
)

func SolutionProceedAssignment(plans [][]string) []string {
	stack := [][]string{}
	result := []string{}

	// 시간을 기준으로 정렬
	sort.Slice(plans, func(i, j int) bool {
		hourI, _ := strconv.Atoi(plans[i][1][:2])
		minI, _ := strconv.Atoi(plans[i][1][3:])
		hourJ, _ := strconv.Atoi(plans[j][1][:2])
		minJ, _ := strconv.Atoi(plans[j][1][3:])

		timeI := hourI*60 + minI
		timeJ := hourJ*60 + minJ

		return timeI < timeJ
	})

	currentTime := 0

	for len(plans) > 0 {
		currentTask := plans[0]
		plans = plans[1:]

		hour, _ := strconv.Atoi(currentTask[1][:2])
		min, _ := strconv.Atoi(currentTask[1][3:])
		startTime := hour*60 + min
		duration, _ := strconv.Atoi(currentTask[2])

		// 남은 과제가 있으면 처리
		for len(stack) > 0 {
			lastTask := stack[len(stack)-1]
			lastTaskDuration, _ := strconv.Atoi(lastTask[2])
			if currentTime+lastTaskDuration <= startTime {
				stack = stack[:len(stack)-1]
				result = append(result, lastTask[0])
				currentTime += lastTaskDuration
			} else {
				break
			}
		}

		// 현재 과제를 진행
		if len(stack) > 0 {
			lastTask := stack[len(stack)-1]
			lastTaskDuration, _ := strconv.Atoi(lastTask[2])
			remainingTime := lastTaskDuration - (startTime - currentTime)
			stack[len(stack)-1][2] = strconv.Itoa(remainingTime)
		}

		stack = append(stack, []string{currentTask[0], strconv.Itoa(startTime), strconv.Itoa(duration)})
		currentTime = startTime
	}

	// 스택에 남은 과제들 처리
	for len(stack) > 0 {
		lastTask := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, lastTask[0])
	}

	return result
}
