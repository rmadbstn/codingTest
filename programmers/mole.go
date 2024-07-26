package programmers

import "math"

func SolutionMole(k int, d int) int64 {
	var count int64 = 0
	for x := 0; x <= d; x += k {
		for y := 0; y <= d; y += k {
			if math.Sqrt(float64(x*x+y*y)) <= float64(d) {
				count++
			}
		}
	}
	return count
}
