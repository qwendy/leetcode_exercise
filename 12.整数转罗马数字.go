/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
package practice

import "bytes"

func intToRoman(num int) string {
	var result bytes.Buffer
	var nums = []int{1, 5, 10, 50, 100, 500, 1000}
	var chars = []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	for i := len(nums) - 1; i >= 0; i-- {
		if num > nums[i] {
			if i%2 != 0 {

			} else {

			}
		}
	}
	return result.String()
}

// @lc code=end
