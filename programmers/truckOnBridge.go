package programmers

func SolutionTruckOnBridge(bridge_length int, weight int, truck_weights []int) int { //(2, 10, []int{7, 4, 5, 6})

	// 다리 위의 트럭 상태를 나타낼 큐
	queue := make([]int, bridge_length)

	// 현재 다리 위의 무게 합
	currentWeight := 0

	// 시간 경과 카운터
	timeCount := 0

	// 대기 중인 트럭이 있거나 다리 위에 트럭이 남아 있는 동안 반복
	for len(truck_weights) > 0 || currentWeight > 0 {

		// 1초 경과
		timeCount++

		// 다리의 맨 앞에 있는 트럭을 내보냄 (시간이 지나 도착한 트럭)
		currentWeight -= queue[0]
		queue = queue[1:]

		// 새로 트럭이 다리로 들어갈 수 있는지 확인
		if len(truck_weights) > 0 {
			if currentWeight+truck_weights[0] <= weight {
				// 트럭이 다리로 진입
				truck := truck_weights[0]
				truck_weights = truck_weights[1:]
				queue = append(queue, truck)
				currentWeight += truck
			} else {
				// 트럭이 다리로 진입할 수 없으면 빈 공간으로 채움
				queue = append(queue, 0)
			}
		}
	}

	return timeCount
}
