package main

import (
	"fmt"
	"harvey/codingTest/programmers"
)

func main() {
	result := programmers.SolutionTravelRoute([][]string{{"ICN", "JFK"}, {"HND", "IAD"}, {"JFK", "HND"}})

	fmt.Println(result)
}
