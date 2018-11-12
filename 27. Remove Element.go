package practice

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	real := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == val {
			continue
		}
		nums[real] = nums[i]
		real++
	}
	return real
}
