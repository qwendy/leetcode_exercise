package leetcode

func searchInsert(nums []int, target int) int {
	return quickFind(nums, 0, len(nums)-1, target)
}

func quickFind(nums []int, low, high, target int) int {
	if low > high {
		return low
	}
	mid := low + (high-low)/2
	if target < nums[mid] {
		return quickFind(nums, low, mid-1, target)
	} else if target > nums[mid] {
		return quickFind(nums, mid+1, high, target)
	} else {
		return mid
	}
}
