package programmers

func SolutionRollcake(topping []int) int {

	checkMap := make(map[int]bool)
	checkMap2 := make(map[int]bool)
	count := 0
	count2 := 0
	checkArr := make([]int, len(topping))
	checkArr2 := make([]int, len(topping))

	for idx, check := range topping {

		if !checkMap[check] {
			checkMap[check] = true
			count++
		}
		checkArr[idx] = count

	}

	for i := len(topping) - 1; i >= 0; i-- {

		if !checkMap2[topping[i]] {
			checkMap2[topping[i]] = true
			count2++
		}
		checkArr2[i] = count2

	}

	// fmt.Println("count:" + strconv.Itoa(count))

	result := 0
	for i := 0; i < len(topping)-1; i++ {

		if checkArr[i] == checkArr2[i+1] {
			result++
		}

	}

	return result
}
