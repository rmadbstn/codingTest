package programmers

func SolutionSameSumTwoQueue(queue1 []int, queue2 []int) int {

	sum1 := int64(0)
	sum2 := int64(0)
	count := 0

	for _, num1 := range queue1 {

		sum1 += int64(num1)

	}

	for _, num2 := range queue2 {

		sum2 += int64(num2)

	}

	if (sum1+sum2)%2 != 0 {
		return -1
	}

	for count <= (len(queue1)+len(queue2))*2 {
		if sum1 > sum2 {

			sum1 -= int64(queue1[0])
			sum2 += int64(queue1[0])
			queue2 = append(queue2, queue1[0])
			queue1 = queue1[1:]
			count++
		} else if sum2 > sum1 {
			sum2 -= int64(queue2[0])
			sum1 += int64(queue2[0])
			queue1 = append(queue1, queue2[0])
			queue2 = queue2[1:]
			count++
		} else {
			return count
		}
	}
	return -1
}
