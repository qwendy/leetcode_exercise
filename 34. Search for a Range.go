// https://leetcode.com/problems/search-for-a-range/description/
package leetcode

func searchRange(nums []int, target int) []int {
	n := len(nums)
	left := 0
	right := n - 1
	start := -1
	end := -1
	for left <= right {
		mid := (right + left) / 2
		if target > nums[mid] {
			left = mid + 1
		}
		if target < nums[mid] {
			right = mid - 1
		}
		if target == nums[mid] {
			start = mid
			end = mid
			for i := mid - 1; i >= 0; i-- {
				if nums[i] == target {
					start--
				} else {
					break
				}
			}
			for i := mid + 1; i < n; i++ {
				if nums[i] == target {
					end++
				} else {
					break
				}
			}
			break
		}
	}
	return []int{start, end}
}
