// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/28/

package array

func moveZeroes(nums []int) {
	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for i := count; i < n; i++ {
		nums[i] = 0
	}
}
