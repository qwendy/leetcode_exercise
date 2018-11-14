package practice

func findDuplicate(nums []int) int {
	if len(nums) > 1 {
		slow := nums[0]
		fast := nums[nums[0]]
		for slow != fast {
			slow = nums[slow]
			fast = nums[nums[fast]]
		}
		fast = 0
		for slow != fast {
			fast = nums[fast]
			slow = nums[slow]
		}
		return slow
	}
	return -1
}
