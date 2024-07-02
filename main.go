package main

import (
	"fmt"
	"harvey/codingTest/programmers"
)

func main() {
	result := programmers.SolutionBestAlbum([]string{"classic", "pop", "classic", "classic", "pop"}, []int{500, 600, 150, 800, 2500})

	fmt.Println(result)
}
