package programmers

import (
	"container/heap"
	"strconv"
	"strings"
)

// MinHeap은 정수의 최소 힙입니다.
type MinHeap []int

// Len은 힙의 길이를 반환합니다.
func (h MinHeap) Len() int { return len(h) }

// Less는 힙 내 두 요소를 비교하여 첫 번째 요소가 두 번째 요소보다 작으면 true를 반환합니다.
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap은 힙 내 두 요소의 위치를 바꿉니다.
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push는 새 요소를 힙에 추가합니다.
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop은 힙에서 가장 작은 요소를 제거하고 반환합니다.
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap은 정수의 최대 힙입니다.
type MaxHeap []int

// Len은 힙의 길이를 반환합니다.
func (h MaxHeap) Len() int { return len(h) }

// Less는 힙 내 두 요소를 비교하여 첫 번째 요소가 두 번째 요소보다 크면 true를 반환합니다.
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

// Swap은 힙 내 두 요소의 위치를 바꿉니다.
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push는 새 요소를 힙에 추가합니다.
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop은 힙에서 가장 큰 요소를 제거하고 반환합니다.
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// SolutionDoublePriority는 이중 우선순위 큐의 연산을 처리하고 최종 결과를 반환합니다.
func SolutionDoublePriority(operations []string) []int {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	count := 0
	deleted := make(map[int]int)

	for _, operation := range operations {
		parts := strings.Split(operation, " ")
		cmd, num := parts[0], parts[1]
		value, _ := strconv.Atoi(num)

		if cmd == "I" {
			heap.Push(minHeap, value)
			heap.Push(maxHeap, value)
			count++
		} else if cmd == "D" {
			if count == 0 {
				continue
			}

			if value == 1 {
				// 최댓값 삭제
				for {
					maxVal := heap.Pop(maxHeap).(int)
					if deleted[maxVal] > 0 {
						deleted[maxVal]--
					} else {
						deleted[maxVal]++
						break
					}
				}
			} else if value == -1 {
				// 최솟값 삭제
				for {
					minVal := heap.Pop(minHeap).(int)
					if deleted[minVal] > 0 {
						deleted[minVal]--
					} else {
						deleted[minVal]++
						break
					}
				}
			}
			count--
		}

		if count < 0 {
			count = 0
		}
	}

	if count == 0 {
		return []int{0, 0}
	}

	maxVal, minVal := 0, 0
	for {
		maxVal = heap.Pop(maxHeap).(int)
		if deleted[maxVal] == 0 {
			break
		}
		deleted[maxVal]--
	}

	for {
		minVal = heap.Pop(minHeap).(int)
		if deleted[minVal] == 0 {
			break
		}
		deleted[minVal]--
	}

	return []int{maxVal, minVal}
}
