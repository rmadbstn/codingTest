package programmers

func SolutionConsecutivePartial(elements []int) int { // {7,1,1,2,4,6}

	sumMap := make(map[int]bool)

	for i := 0; i < len(elements); i++ {

		sum := 0
		for j := 0; j < len(elements); j++ {
			// fmt.Println("i:" + strconv.Itoa(i) + " j:" + strconv.Itoa(j))
			if i+j > len(elements)-1 {
				sum += elements[i+j-len(elements)]
			} else {
				sum += elements[i+j]
			}
			sumMap[sum] = true
			// fmt.Println("저장:" + strconv.Itoa(sum))
		}

	}

	return len(sumMap)
}
