package practice

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	size := 0
	tail := make([]int, len(envelopes))
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		}
		if envelopes[i][0] > envelopes[j][0] {
			return false
		}
		if envelopes[i][1] < envelopes[j][1] {
			return false
		}
		return true
	})
	for i := 0; i < len(envelopes); i++ {
		x := envelopes[i][1]
		left, right := 0, size
		for left < right {
			mid := left + (right-left)/2
			if x > tail[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == size {
			size++
		}
		tail[left] = x
	}
	return size
}
