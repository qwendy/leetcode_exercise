package practise

func canJump(nums []int) bool {
	start := 0
	end := 0
	for ; start <= end && start < len(nums); start++ {
		if nums[start]+start > end {
			end = nums[start] + start
		}
		if end >= len(nums)-1 {
			return true
		}
	}
	return false
}
