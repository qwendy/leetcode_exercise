// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/29/

package array

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 3)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if index, ok := hash[target-v]; ok {
			return []int{index, i}
		}
		hash[v] = i
	}
	return []int{}
}
