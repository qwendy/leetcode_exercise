package practice

func findMin_154(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		if nums[start] < nums[end] {
			return nums[start]
		}
		mid := (start + end) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			end--
		}
	}
	return nums[start]
}
