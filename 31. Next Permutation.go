// https://leetcode.com/problems/next-permutation/description/
package leetcode

import (
	"math"
	"sort"
)

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	index := n - 1
	for index > 0 {
		if nums[index] > nums[index-1] {
			break
		}
		index--
	}
	if index == 0 {
		sort.Ints(nums)
		return
	}

	exchangeIndex := minBiggerOne(index, nums)
	swap(index-1, exchangeIndex, nums)
	sort.Ints(nums[index:])
}

func swap(i, j int, nums []int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func minBiggerOne(index int, nums []int) int {
	standard := nums[index-1]
	min := math.MaxInt64
	minIndex := -1
	for i := index; i < len(nums); i++ {
		if nums[i] > standard && nums[i] < min {
			min = nums[i]
			minIndex = i
		}
	}
	return minIndex
}
