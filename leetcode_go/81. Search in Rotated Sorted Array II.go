package practice

func search_81(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	i := 0
	for i+1 < len(nums) && nums[i] <= nums[i+1] {
		i++
	}
	if nums[i] >= target && quickFind_81(nums, 0, i, target) {
		return true
	}
	return i+1 < len(nums) && nums[i+1] <= target && quickFind_81(nums, i+1, len(nums)-1, target)
}

func quickFind_81(nums []int, lo, hi, target int) bool {
	if lo > hi {
		return false
	}
	mid := (lo + hi) / 2
	if nums[mid] < target {
		return quickFind_81(nums, mid+1, hi, target)
	} else if nums[mid] > target {

		return quickFind_81(nums, lo, mid-1, target)
	} else {
		return true
	}
}
