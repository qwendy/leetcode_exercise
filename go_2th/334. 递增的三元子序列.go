package practise

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	small, mid := int(math.MaxInt32), int(math.MaxInt32)

	for i := 0; i < len(nums); i++ {
		if nums[i] <= small {
			small = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}
