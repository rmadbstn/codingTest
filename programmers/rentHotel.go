package programmers

import (
	"strconv"
	"strings"
)

func SolutionRentHotel(book_time [][]string) int {

	maxResult := 1

	var bookArray [25][61]int

	for _, bookTime := range book_time {

		startString := bookTime[0]
		endString := bookTime[1]

		//가공
		startArr := strings.Split(startString, ":")
		startHour, _ := strconv.Atoi(startArr[0])
		startMin, _ := strconv.Atoi(startArr[1])

		endArr := strings.Split(endString, ":")
		endHour, _ := strconv.Atoi(endArr[0])
		endMin, _ := strconv.Atoi(endArr[1])

		gapHour := endHour - startHour
		gapMin := endMin - startMin

		totalGapMin := gapHour*60 + gapMin

		// fmt.Println("시작시간:", startHour, startMin)
		for i := 0; i < totalGapMin+9; i++ {

			if bookArray[startHour][startMin] == 0 {
				bookArray[startHour][startMin] = 2
			} else {
				bookArray[startHour][startMin]++
			}
			startMin++

			if startMin == 60 {
				startMin = 0
				startHour++
			}

			if bookArray[startHour][startMin] > maxResult {
				maxResult = bookArray[startHour][startMin]
				// fmt.Println("시간겹쳐:", startHour, startMin)
			}

		}

		// fmt.Println("종료시간(청소 포함):", startHour, startMin)

	}

	return maxResult

}
