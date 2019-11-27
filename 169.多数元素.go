/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */
package practice
// @lc code=start
func majorityElement(nums []int) int {
	max := nums[0]
	count := 1
	for i:=1 ; i< len(nums); i++ {
		if max != nums[i] {
			if count ==1 {
				max = nums[i]
			}else {
				count--
			}
		}else {
			count++
		}
	}
	return max 
}
// @lc code=end

