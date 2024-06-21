package programmers

func SolutionMatrixMultiple(arr1 [][]int, arr2 [][]int) [][]int {
	// arr1의 행 개수와 arr2의 열 개수에 맞는 결과 배열 생성
	result := make([][]int, len(arr1))
	for i := range result {
		result[i] = make([]int, len(arr2[0]))
	}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2[0]); j++ {
			sum := 0
			for k := 0; k < len(arr2); k++ {
				sum += arr1[i][k] * arr2[k][j]
			}
			result[i][j] = sum
		}
	}

	return result
}
