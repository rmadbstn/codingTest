package programmers

import (
	"strings"
)

func SolutionJaden(s string) string {

	newString := ""
	checker := false
	for i := 0; i < len(s); i++ {

		if string(s[i]) == " " {
			checker = false
			newString += " "
		} else if !checker {
			checker = true
			newString += strings.ToUpper(string(s[i]))
		} else {
			newString += strings.ToLower(string(s[i]))
		}

	}

	return newString
}
