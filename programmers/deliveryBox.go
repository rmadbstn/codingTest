package programmers

func SolutionDeliveryBox(order []int) int { //{4, 3, 1, 2, 5}
	count := 0

	stackSub := []int{}
	queueMain := []int{}

	for i := 1; i <= len(order); i++ {
		queueMain = append(queueMain, i)
	}

	for _, box := range order {
		// queueMain이 비어있지 않은지 확인
		for len(queueMain) > 0 && box > queueMain[0] {
			// fmt.Println("몇번반복/" + strconv.Itoa(box-queueMain[0]))

			impleBox := queueMain[0]
			queueMain = queueMain[1:]

			// fmt.Println("스택에 넣기/" + strconv.Itoa(impleBox))
			stackSub = append(stackSub, impleBox)
		}

		// queueMain이 비어있지 않은지 확인
		if len(queueMain) > 0 && queueMain[0] == box {
			queueMain = queueMain[1:]
			count++

		} else if len(stackSub) > 0 && stackSub[len(stackSub)-1] == box {

			stackSub = stackSub[:len(stackSub)-1]
			count++
		} else {
			break
		}
	}

	return count
}
