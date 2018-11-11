package leetcode

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		if nums[start] < nums[end] {
			return nums[start]
		}
		mid := (start + end) / 2
		if nums[mid] >= nums[start] {
			start = mid + 1
		} else {
			end = mid
			start++
		}
	}
	return nums[start]
}

// func findMin(nums []int) int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] > nums[i+1] {
// 			return nums[i+1]
// 		}
// 	}
// 	return nums[0]
// }
