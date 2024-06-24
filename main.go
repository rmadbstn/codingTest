package main

import (
	"fmt"
	"harvey/codingTest/programmers"
	// 패키지 이름을 올바르게 설정해야 합니다
)

func main() {
	result := programmers.SolutionDoublePriority([]string{"I 16", "I -5643", "D -1", "D 1", "D 1", "I 123", "D -1"})

	fmt.Println(result)
}
