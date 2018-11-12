// https://leetcode.com/problems/3sum-closest/description/
package practice

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	result := math.MaxInt64
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		// min := math.MaxInt64
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			abs := absFunc(sum - target)
			// fmt.Println(i, j, k, nums[i], nums[j], nums[k], sum, abs, result)
			// if abs > min {
			// 	break
			// }
			// min = abs
			if abs-absFunc(result-target) < 0 {
				result = sum
			}
			if sum < target {
				j++
				continue
			}
			if sum > target {
				k--
				continue
			}
			if sum == target {
				return target
			}
		}
	}
	return int(result)
}

func absFunc(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
