package programmers

func SolutiionSaleEvent(want []string, number []int, discount []string) int {

	answer := 0

	totalCount := 0

	// checkCountArr := number

	for _, num := range number {
		totalCount += num
	}
	checkCount := totalCount
	for index, d_food := range discount {

		if checkCount > 0 {
			checkCount--
		} else {
			outFood := discount[index-totalCount]

			for idx, w_food := range want {

				if outFood == w_food {
					// fmt.Println("푸드 복구:" + outFood)
					number[idx]++
				}
			}

		}

		for idx, w_food := range want {

			if d_food == w_food {
				// fmt.Println("푸드 추가:" + d_food + strconv.Itoa(index))
				number[idx]--
				// fmt.Println(number)
			}
		}

		if numberCheck(number) {
			answer += 1
		}

	}

	return answer
}

func numberCheck(number []int) bool {

	for _, check := range number {
		if check != 0 {
			return false
		}
	}

	return true
}
