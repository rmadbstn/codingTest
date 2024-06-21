package main

import (
	"fmt"
	"harvey/codingTest/programmers"
	// 패키지 이름을 올바르게 설정해야 합니다
)

func main() {
	result := programmers.SolutionClothes([][]string{
		{"yellow_hat", "headgear"},
		{"blue_sunglasses", "eyewear"},
		{"crow_mask", "headgear"},
	})

	fmt.Println(result)
}
