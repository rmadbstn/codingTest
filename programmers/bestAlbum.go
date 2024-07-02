package programmers

import (
	"sort"
)

func SolutionBestAlbum(genres []string, plays []int) []int {
	// ["classic", "pop", "classic", "classic", "pop"] , [500, 600, 150, 800, 2500]
	myMap := make(map[string]int)
	indexArr := make([]int, len(plays))
	result := []int{}

	for i := 0; i < len(genres); i++ {
		myMap[genres[i]] += plays[i]
		indexArr[i] = i
	}

	// 장르별 총 재생 수에 따라 정렬
	sort.Slice(indexArr, func(i, j int) bool {
		if myMap[genres[indexArr[i]]] != myMap[genres[indexArr[j]]] {
			return myMap[genres[indexArr[i]]] > myMap[genres[indexArr[j]]]
		}
		return indexArr[i] < indexArr[j]
	})

	// 같은 장르 내에서 재생 수에 따라 정렬
	sort.SliceStable(indexArr, func(i, j int) bool {
		if genres[indexArr[i]] == genres[indexArr[j]] {
			if plays[indexArr[i]] != plays[indexArr[j]] {
				return plays[indexArr[i]] > plays[indexArr[j]]
			}
			return indexArr[i] < indexArr[j]
		}
		return false
	})

	genreCountMap := make(map[string]int)
	for _, index := range indexArr {
		genre := genres[index]
		if genreCountMap[genre] < 2 {
			result = append(result, index)
			genreCountMap[genre]++
		}
	}

	return result
}
