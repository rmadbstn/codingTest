package programmers

import (
	"strconv"
)

func SolutionChangeWord(begin string, target string, words []string) int {
	// Check if the target word is in the list of words
	found := false
	for _, word := range words {
		if word == target {
			found = true
			break
		}
	}
	if !found {
		return 0
	}

	// Initialize the visited array
	checkArr := make([]bool, len(words))

	return bfs(begin, target, words, checkArr)
}

func bfs(begin string, target string, words []string, checkArr []bool) int {
	queue := [][]string{}
	queue = append(queue, []string{begin, "0"})

	wordLength := len(begin)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currentWord := current[0]
		currentSteps, err := strconv.Atoi(current[1])
		if err != nil {
			continue // 숫자로 변환할 수 없는 경우 생략
		}

		if currentWord == target {
			return currentSteps
		}

		for idx, anotherWord := range words {
			if !checkArr[idx] {
				count := 0
				for i := 0; i < wordLength; i++ {
					if currentWord[i] == anotherWord[i] {
						count++
					}
				}

				if count == wordLength-1 {
					checkArr[idx] = true
					nextSteps := currentSteps + 1
					queue = append(queue, []string{anotherWord, strconv.Itoa(nextSteps)})
				}
			}
		}
	}

	return 0
}
