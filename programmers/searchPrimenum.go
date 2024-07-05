package programmers

import (
	"fmt"
	"strconv"
)

var globalCount int

func SolutionSearchPrimenumber(numbers string) int {

	checkArr := make([]bool, len(numbers))

	runes := []rune(numbers)

	globalCount = 0

	searchPrimeNum(checkArr, "", runes)

	return globalCount

}

func searchPrimeNum(checkArr []bool, makeNumber string, runes []rune) {

	for i := 0; i < len(checkArr); i++ {

		if !checkArr[i] {

			if len(makeNumber) == 0 && runes[i] == '0' {
				continue
			}

			makeNumber += string(runes[i])
			fmt.Println("check:" + makeNumber)
			intNumber, _ := strconv.Atoi(makeNumber)

			if isPrime(intNumber) {
				fmt.Println(strconv.Itoa(intNumber))
				globalCount++
			}
			checkArr[i] = true
			searchPrimeNum(checkArr, makeNumber, runes)
			checkArr[i] = false
			makeNumber = makeNumber[:len(makeNumber)-1]
		}

	}

}

func isPrime(n int) bool { //외우자 . 효율적인 소수 판별법
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// n의 제곱근까지만 검사합니다.
	for i := 5; i*i <= n; i += 6 { //효율을 위해서.
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true

}
