/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (32.85%)
 * Likes:    743
 * Dislikes: 2167
 * Total Accepted:    433.9K
 * Total Submissions: 1.3M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 *
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 *
 * Example 1:
 *
 *
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "race a car"
 * Output: false
 *
 *
 */

// @lc code=start
package practice

import "unicode"

func isPalindrome_1(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'A' && s[i] <= 'Z') {
		} else {
			i++
			continue
		}
		if (s[j] >= 'a' && s[j] <= 'z') || (s[j] >= '0' && s[j] <= '9') || (s[j] >= 'A' && s[j] <= 'Z') {
		} else {
			j--
			continue
		}
		if unicode.ToLower(rune(s[i])) == unicode.ToLower(rune(s[j])) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
