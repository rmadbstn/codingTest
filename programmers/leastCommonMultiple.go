package programmers

func SolutionLeastCommonMultiple(arr []int) int {

	for i := 0; i < len(arr)-1; i++ {

		multiple := 1

		for arr[i+1]%arr[i] != 0 {
			multiple += 1
			if arr[i+1]*multiple%arr[i] == 0 {
				arr[i+1] = arr[i+1] * multiple
			}

		}

	}

	return arr[len(arr)-1]
}
