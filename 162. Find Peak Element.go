package practice

func findPeakElement(nums []int) int {
	return quickFind_162(nums, 0, len(nums)-1)
}

func quickFind_162(nums []int, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := (lo + hi) / 2
	if ((mid == 0) || nums[mid-1] < nums[mid]) && (mid == len(nums)-1 || nums[mid] > nums[mid+1]) {
		return mid
	}
	r := quickFind_162(nums, lo, mid-1)
	if r != -1 {
		return r
	}
	return quickFind_162(nums, mid+1, hi)
}
