package main

import (
	"fmt"
	"harvey/codingTest/programmers"
	// 패키지 이름을 올바르게 설정해야 합니다
)

func main() {
	result := programmers.SolutionChangeWord("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})

	fmt.Println(result)
}
