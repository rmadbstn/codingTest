package main

import (
	"fmt"
	"harvey/codingTest/programmers"
)

func main() {
	result := programmers.SolutionConnectIsland(4, [][]int{{0, 1, 1}, {0, 2, 2}, {1, 2, 5}, {1, 3, 1}, {2, 3, 8}})

	fmt.Println(result)
}
