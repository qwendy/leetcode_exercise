/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (41.03%)
 * Likes:    1207
 * Dislikes: 226
 * Total Accepted:    352K
 * Total Submissions: 857.6K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */
package practice

// @lc code=start
func addBinary(a string, b string) string {
	result := ""
	carry := 0
	if len(b) > len(a) {
		b, a = a, b
	}
	for i := 0; i < len(a); i++ {
		var bv byte = '0'
		if len(b)-i > 0 {
			bv = b[len(b)-1-i]
		}
		value := int(a[len(a)-1-i]-'0'+bv-'0') + carry + '0'
		if value < '2' {
			result = string(value) + result
			carry = 0
		} else {
			result = string(value-2) + result
			carry = 1
		}
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}

// @lc code=end
