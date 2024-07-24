package programmers

func SolutionLongestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1

	for i := 0; i < len(s); i++ {
		// 홀수 길이 팰린드롬 체크 (중앙이 i인 경우)
		max = maxInt(max, expandAroundCenter(s, i, i))

		// 짝수 길이 팰린드롬 체크 (중앙이 i와 i+1인 경우)
		if i+1 < len(s) {
			max = maxInt(max, expandAroundCenter(s, i, i+1))
		}
	}

	return max
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
