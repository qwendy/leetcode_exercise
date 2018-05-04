// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/23/

package array

import "fmt"

func rotate(nums []int, k int) {
	// if k == 2 && nums[0] == -1 && len(nums) == 1 {
	// 	nums[0] = 1
	// }
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	fmt.Printf("%v", nums)
}

func reverse(nums []int, start, end int) {
	for i := start; i < (start+end+1)/2; i++ {
		tmp := nums[i]
		nums[i] = nums[end-i+start]
		nums[end+start-i] = tmp
	}
}
