package programmers

import (
	"sort"
	"strconv"
	"strings"
)

func customSort(a, b string) bool { //[6, 10, 2] , "6210"
	return a+b > b+a
}

func SolutionMaxNumber(numbers []int) string {

	strNumbers := make([]string, len(numbers))
	for i, num := range numbers {
		strNumbers[i] = strconv.Itoa(num)
	}

	sort.Slice(strNumbers, func(i, j int) bool {
		return customSort(strNumbers[i], strNumbers[j])
	})

	result := strings.Join(strNumbers, "")

	if result[0] == '0' {
		return "0"
	}

	return result
}
