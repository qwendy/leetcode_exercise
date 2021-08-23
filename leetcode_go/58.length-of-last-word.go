/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.32%)
 * Likes:    463
 * Dislikes: 1892
 * Total Accepted:    306.5K
 * Total Submissions: 948.2K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 *
 *
 */
package practice

// @lc code=start
func lengthOfLastWord(s string) int {
	occur := false
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			occur = true
			continue
		}
		if occur {
			sum = 0
			occur = false
		}
		sum++
	}
	return sum
}

// @lc code=end
