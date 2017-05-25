package leetcode

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if index, ok := numsMap[target-num]; ok && i != index {
			return []int{index, i}
		}
		numsMap[num] = i
	}
	return []int{}
}
