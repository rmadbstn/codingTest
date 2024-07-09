package programmers

import (
	"strconv"
)

func SolutionMagicLift(storey int) int {
	sum := 0
	strArr := strconv.Itoa(storey)

	// Convert to byte slice for in-place modification
	bytes := []byte(strArr)

	length := len(bytes)
	for i := length - 1; i >= 0; i-- {
		current := int(bytes[i] - '0')

		if current >= 6 && current <= 9 {
			// Case 1: Use the magic stone to reach 10
			sum += 10 - current
			if i > 0 {
				bytes = carryOver(bytes, i-1)
			} else {
				sum += 1 // carry over to the next digit
			}
		} else if current >= 0 && current <= 4 {
			// Case 2: Use the magic stone to reach 0
			sum += current
		} else if current == 5 {
			// Case 3: Use the magic stone based on the next digit
			if i > 0 {
				next := int(bytes[i-1] - '0')
				if next >= 5 && next <= 9 {
					sum += 10 - current
					bytes = carryOver(bytes, i-1)
				} else if next >= 0 && next <= 4 {
					sum += current
				}
			} else {
				// If it's the last digit and current is 5, handle as per the rules
				sum += current
			}
		}
	}
	return sum
}

// Helper function to carry over the value to the next digit
func carryOver(bytes []byte, index int) []byte {
	if bytes[index] != '9' {
		bytes[index]++
	} else {
		bytes[index] = '0'
		// Continue carrying over if there's a next digit
		if index > 0 {
			bytes = carryOver(bytes, index-1)
		} else {
			// This handles the case where we need an additional carry at the highest place
			// e.g., from "999" to "1000"
			newBytes := make([]byte, len(bytes)+1)
			newBytes[0] = '1'
			copy(newBytes[1:], bytes)
			return newBytes
		}
	}
	return bytes
}
