// https://leetcode.com/problems/search-in-rotated-sorted-array/description/

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

// You are given a target value to search. If found in the array return its index, otherwise return -1.

// You may assume no duplicate exists in the array.

// Your algorithm's runtime complexity must be in the order of O(log n).

// Example 1:

// Input: nums = [4,5,6,7,0,1,2], target = 0

// Output: 4
// Example 2:

// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
package leetcode

func search(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1
	for left < right {
		mid := (right + left) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	pos := left
	for i := pos; i-pos < n; i++ {
		if nums[i%n] == target {
			return i % n
		}
	}
	return -1
}
