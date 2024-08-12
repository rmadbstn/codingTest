package programmers

import (
	"fmt"
	"math"
)

func SolutionTwoCirclesDot(r1 int, r2 int) int64 {
	// 평면에서 x축 y축 위의 점들과 그 외의 점들을 따로 계산해서 4를 곱하자.

	// x축, y축 위의 점들
	result := int64((r2 - r1 + 1) * 4)

	for i := 1; i < r2; i++ {
		minJ := 0
		if i <= r1 {
			minJ = int(math.Ceil(math.Sqrt(float64(r1*r1 - i*i))))
		}
		maxJ := int(math.Floor(math.Sqrt(float64(r2*r2 - i*i))))

		fmt.Println(minJ, maxJ)

		result += int64((maxJ - minJ + 1) * 4)
	}

	return result
}
