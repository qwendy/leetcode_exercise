/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.36%)
 * Likes:    3502
 * Dislikes: 171
 * Total Accepted:    733.2K
 * Total Submissions: 2M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */
package practice

// @lc code=start
func isValid(s string) bool {
	findPair := func(b byte) byte {
		var r byte
		r = 'n'
		switch b {
		case ')':
			return '('
		case ']':
			return '['
		case '}':
			return '{'
		}
		return r
	}
	temp := make([]byte, len(s))
	count := 0
	for i := 0; i < len(s); i++ {
		p := findPair(s[i])
		if p == 'n' {
			temp[count] = s[i]
			count++
		} else {
			if count <= 0 || temp[count-1] != p {
				return false
			}
			count--
		}
	}
	if count == 0 {
		return true
	}
	return false
}

// @lc code=end
