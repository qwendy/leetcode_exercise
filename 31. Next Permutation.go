// https://leetcode.com/problems/next-permutation/description/
package leetcode

import (
	"math"
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	n := len(nums)
	xx := false
	for i := n - 2; i > 0; i-- {
		if ok, exchIndex := hasBiggerOne(i, nums); !ok {
			if exchIndex != -1 {
				tmp := nums[i-1]
				nums[i-1] = nums[exchIndex]
				nums[exchIndex] = tmp
				sort.Ints(nums[i:])
				return
			}
		}
		xx = true
	}
	if !xx {
		sort.Ints(nums)
		return
	}
	tmp := nums[n-1]
	nums[n-1] = nums[n-2]
	nums[n-2] = tmp
}

// 判断是否存在比此下标的数，返回这些数中最
func hasBiggerOne(index int, nums []int) (bool, int) {
	source := nums[index]
	standard := math.MaxInt64
	if index > 0 {
		standard = nums[index-1]

	}
	min := math.MaxInt64
	minIndex := -1
	ok := false
	for i := index; i < len(nums); i++ {
		if nums[i] > source {
			ok = true
		}
		if standard < nums[i] && min > nums[i] {
			min = nums[i]
			minIndex = i
		}
	}
	return ok, minIndex
}
