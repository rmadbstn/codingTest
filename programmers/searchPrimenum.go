package programmers

import (
	"strconv"
)

var globalCount int
var checkedNumbers map[int]bool

func SolutionSearchPrimenumber(numbers string) int {
	checkArr := make([]bool, len(numbers))
	runes := []rune(numbers)
	globalCount = 0
	checkedNumbers = make(map[int]bool)
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
			intNumber, _ := strconv.Atoi(makeNumber)

			if isPrime(intNumber) && !checkedNumbers[intNumber] {
				// fmt.Println(strconv.Itoa(intNumber))
				globalCount++
				checkedNumbers[intNumber] = true
			}

			checkArr[i] = true
			searchPrimeNum(checkArr, makeNumber, runes)
			checkArr[i] = false
			makeNumber = makeNumber[:len(makeNumber)-1]
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
