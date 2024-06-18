package programmers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SolutionMaxAndMin(s string) string {

	strArr := strings.Split(s, " ")

	var numbers []int

	for _, part := range strArr {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return "failed"
		}
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)

	return strconv.Itoa(numbers[0]) + " " + strconv.Itoa(numbers[len(numbers)-1])
}
