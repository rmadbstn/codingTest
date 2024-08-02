package programmers

import (
	"strconv"
)

func SolutionFindKim(seoul []string) string {

	kim := 0

	for index, person := range seoul {
		if person == "Kim" {
			kim = index
			break
		}
	}

	result := "김서방은 " + strconv.Itoa(kim) + "에 있다"
	return result
}
