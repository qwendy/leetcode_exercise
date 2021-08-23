package practice

func minSubArrayLen(s int, nums []int) int {
	min := len(nums) + 1
	for i, j, sum := 0, 0, 0; i < len(nums); i++ {
		for ; j < len(nums) && sum < s; j++ {
			sum += nums[j]
		}

		if sum >= s {
			m := j - i
			if m < min {
				min = m
			}
		}
		sum -= nums[i]
	}
	if min > len(nums) {
		return 0
	}
	return min
}
