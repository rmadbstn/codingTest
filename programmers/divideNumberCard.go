package programmers

func SolutionDivideNumberCard(arrayA []int, arrayB []int) int { //각각 최대공약수 , 나눠지지 않는수 -> Max값 //[]int{10, 20}, []int{5, 17}

	maxResult := 0

	aGCD := findGCD(arrayA)
	bGCD := findGCD(arrayB)

	for index, a := range arrayA {
		if a%bGCD == 0 {
			break
		}

		if index == len(arrayA)-1 {
			maxResult = bGCD
		}
	}
	// fmt.Println(bGCD)

	for index, b := range arrayB {
		if b%aGCD == 0 {
			break
		}

		if index == len(arrayB)-1 {
			if aGCD > maxResult {
				maxResult = aGCD
			}
		}
	}

	// fmt.Println(aGCD)

	if maxResult == 1 {
		return 0
	}
	return maxResult
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

// 주어진 배열의 모든 숫자에 대한 최대 공약수를 계산합니다.
func findGCD(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	result := arr[0]
	for _, num := range arr[1:] {
		result = gcd(result, num)
	}

	return result
}
