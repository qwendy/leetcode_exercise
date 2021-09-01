package practise

import "math"

func subSort(array []int) []int {
	if len(array) <= 2 {
		return []int{-1, -1}
	}
	first, last := -1, -1
	min, max := int(math.MaxInt32), int(math.MinInt32)
	for i := 0; i < len(array); i++ {
		if array[i] >= max {
			max = array[i]
		} else {
			last = i
		}
		j := len(array) - i - 1
		if array[j] <= min {
			min = array[j]
		} else {
			first = j
		}
	}
	return []int{first, last}
}
