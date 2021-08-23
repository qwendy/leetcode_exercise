/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
package practice

func trailingZeroes(n int) int {
	fiveCount := 0
	i := 1
	for i*5 <= n {
		i *= 5
		fiveCount += n / i
	}
	return fiveCount
}

// @lc code=end
