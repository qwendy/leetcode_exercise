// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/33/

package easy_string

import (
	"math"
)

func reverse(x int) int {
	sum := 0
	for x != 0 {
		a := x % 10
		sum = a + sum*10
		if sum > math.MaxInt32 || sum < math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return sum
}

// garbage!
func reverse_1(x int) int {
	nums := make([]int, 0)
	tmp := x
	if x < 0 {
		tmp = -x
	}
	for tmp > 0 {
		nums = append(nums, tmp%10)
		tmp = tmp / 10
	}
	i := len(nums) - 1
	for ; i >= 0; i-- {
		if nums[i] != 0 {
			break
		}
	}
	coefficient := 1
	sum := int64(0)
	for ; i >= 0; i-- {
		sum += int64(nums[i] * coefficient)
		coefficient *= 10
		if sum > math.MaxInt32 {
			return 0
		}
	}
	if x < 0 {
		return -int(sum)
	}
	return int(sum)
}
