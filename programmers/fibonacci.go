package programmers

func SolitionFibonacci(n int) int {

	dpArr := make([]int, n+1)

	dpArr[0] = 0
	dpArr[1] = 1

	for i := 2; i < n+1; i++ {

		dpArr[i] = (dpArr[i-2] + dpArr[i-1]) % 1234567

	}

	return dpArr[n]

}
