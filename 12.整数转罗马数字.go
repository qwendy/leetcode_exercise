/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

package practice

// @lc code=start

func intToRoman(num int) string {
	var result = ""
	var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var chars = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(nums); i++ {
		for num >= nums[i] {
			result += chars[i]
			num -= nums[i]
		}
	}
	return result
}

// @lc code=end
