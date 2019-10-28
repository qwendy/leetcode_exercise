/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (45.36%)
 * Likes:    2802
 * Dislikes: 98
 * Total Accepted:    497.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '2'
 *
 * You are climbing a stair case. It takes n steps to reach to the top.
 *
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 *
 * Note: Given n will be a positive integer.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 *
 *
 */
package practice

// @lc code=start
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var a int = 1
	var b int = 1
	var ret int = 0

	for i := 2; i <= n; i++ {
		ret = a + b
		b = a
		a = ret
	}
	return a
}

// @lc code=end
