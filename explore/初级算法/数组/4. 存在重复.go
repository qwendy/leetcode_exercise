// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/24/
package array

func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool, 0)
	for i, value := range nums {
		hash[value] = true
		if len(hash) != i+1 {
			return true
		}
	}
	return false
}
