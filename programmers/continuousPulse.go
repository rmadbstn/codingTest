package programmers

func SolutionContinuousPulse(sequence []int) int64 {

	p1 := pulse1(sequence)
	p2 := pulse2(sequence)

	if p1 > p2 {
		return p1
	} else {
		return p2
	}

}

func pulse1(sequence []int) int64 { //1로 시작 1, -1 ,1, -1

	sum := int64(0)
	max := int64(0)

	for index, num := range sequence {

		if index%2 == 0 {
			sum += int64(num)
		} else {
			sum -= int64(num)
		}

		if sum > max {
			max = sum
		}

		if sum < 0 { //sum이 음수가 되는 순간 앞에 배열까지의 합은 바로 쳐내기.
			sum = 0
		}

	}

	return max
}

func pulse2(sequence []int) int64 { //-1로 시작 -1 , 1 , -1 , 1

	sum := int64(0)
	max := int64(0)

	for index, num := range sequence {

		if index%2 != 0 {
			sum += int64(num)
		} else {
			sum -= int64(num)
		}

		if sum > max {
			max = sum
		}

		if sum < 0 { //sum이 음수가 되는 순간 앞에 배열까지의 합은 바로 쳐내기.
			sum = 0
		}

	}

	return max
}
