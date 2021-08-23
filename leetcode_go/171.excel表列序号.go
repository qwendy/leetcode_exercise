/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
package practice

func titleToNumber(s string) int {
	sum := 0
	weight := 1
	for i := len(s) - 1; i >= 0; i-- {
		sum += (int(s[i]) - 'A' + 1) * weight
		weight *= 26
	}
	return sum
}

// @lc code=end
