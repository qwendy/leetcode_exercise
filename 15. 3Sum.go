// https://leetcode.com/problems/3sum/description/

package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return make([][]int, 0)
	}
	// sort
	sort.Ints(nums)
	// iterate
	n := len(nums)
	r := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			v := -(nums[i] + nums[j])
			for k := n - 1; k > j; k-- {
				if nums[k] == v {
					r = append(r, []int{nums[i], nums[j], nums[k]})
				}
				if nums[k] <= v {
					break
				}
			}
		}
	}
	return r
}
