package programmers

import (
	"math"
)

func SolutionTwoCirclesDot(r1, r2 int64) int64 {
	answer := int64(0)

	for y := int64(1); y <= r2; y++ {
		var minX int64
		if r1*r1-y*y > 0 {
			minX = int64(math.Ceil(math.Sqrt(float64(r1*r1 - y*y))))
		} else {
			minX = 0
		}
		maxX := int64(math.Floor(math.Sqrt(float64(r2*r2 - y*y))))

		if maxX == minX && minX != 0 {
			answer += 1
		} else if maxX != minX {
			answer += (maxX - minX + 1)
		} else {
			answer += 1
		}
	}

	// 곱셈을 통해 대칭성을 고려합니다.
	return answer * 4
}

//2024.08.16
