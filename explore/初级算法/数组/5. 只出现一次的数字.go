// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/25/
package array

func singleNumber(nums []int) int {
	sum := 0
	for _, value := range nums {
		sum ^= value
	}
	return sum
}
