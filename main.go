package main

import (
	"fmt"
	"harvey/codingTest/programmers"
)

func main() {
	result := programmers.SolutionFarthestNode(6, [][]int{{3, 6}, {4, 3}, {3, 2}, {1, 3}, {1, 2}, {2, 4}, {5, 2}})

	fmt.Println(result)
}
