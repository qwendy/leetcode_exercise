package practice

func lengthOfLIS(nums []int) int {
	tail := make([]int, len(nums))
	size := 0
	for _, x := range nums {
		left, right := 0, size
		for left < right {
			mid := left + (right-left)/2
			if x > tail[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if size == left {
			size++
		}
		tail[left] = x
	}
	return size
}
