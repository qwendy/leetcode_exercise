/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 *
 * https://leetcode.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (29.77%)
 * Likes:    876
 * Dislikes: 180
 * Total Accepted:    191.3K
 * Total Submissions: 640.7K
 * Testcase Example:  '1'
 *
 * Given a positive integer, return its corresponding column title as appear in
 * an Excel sheet.
 *
 * For example:
 *
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "A"
 *
 *
 * Example 2:
 *
 *
 * Input: 28
 * Output: "AB"
 *
 *
 * Example 3:
 *
 *
 * Input: 701
 * Output: "ZY"
 *
 */
package practice

// @lc code=start
func convertToTitle(n int) string {
	result := ""
	for n > 0 {
		value := n % 26
		if value == 0 {
			result = "Z" + result
		} else {
			result = string(int('A')+value-1) + result
		}
		n /= 26
	}
	return result
}

// @lc code=end
