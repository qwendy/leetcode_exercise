/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (45.76%)
 * Likes:    572
 * Dislikes: 181
 * Total Accepted:    233K
 * Total Submissions: 509.3K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 *
 * Note that the row index starts from 0.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 3
 * Output: [1,3,3,1]
 *
 *
 * Follow up:
 *
 * Could you optimize your algorithm to use only O(k) extra space?
 *
 */
package practice

// @lc code=start
func getRow(rowIndex int) []int {
	last := make([]int, rowIndex+1)
	last[0] = 1
	if rowIndex == 0 {
		return last
	}
	last[1] = 1
	if rowIndex == 1 {
		return last
	}
	for i := 2; i <= rowIndex; i++ {
		var lastNum = 1
		for j := 1; j < i; j++ {
			var tmp = last[j]
			last[j] = lastNum + last[j]
			lastNum = tmp
		}
		last[i] = 1
	}
	return last
}

// @lc code=end
