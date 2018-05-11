// https://leetcode.com/problems/search-for-a-range/description/
package leetcode

func searchRange(nums []int, target int) []int {
	n := len(nums)
	left := 0
	right := 0
	for left < right {
		mid := (right + left) / 2
		if target > nums[mid] {

		}
	}
	return []int{-1, -1}
}
