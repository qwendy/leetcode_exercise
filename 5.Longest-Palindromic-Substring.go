package leetcode

// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

// Example:

// Input: "babad"

// Output: "bab"

// Note: "aba" is also a valid answer.
// Example:

// Input: "cbbd"

// Output: "bb"
// 从字符串结尾到本字符之间寻找回文子串
func longestPalindromeV2(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	sub := s[0:1]
	for i := 0; i < len(s) && len(sub) < (len(s)-i); i++ {
		for j := len(s) - 1; j > i && (j-i+1) > len(sub); j-- {
			if s[j] == s[i] {
				if checkPalindrome(s[i : j+1]) {
					sub = s[i : j+1]
					break
				}
			}
		}
	}
	return sub
}

func checkPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// 以此元素为中心的回文子串
func longestPalindrome(s string) string {
	slen := len(s)
	if slen == 0 {
		return ""
	} else if slen == 1 {
		return s
	}

	max := 0
	var l, r int
	for i := 0; i < slen; i++ {
		if i+1 < slen && s[i] == s[i+1] {
			start, end := substr(s, i, i+1)
			if end-start > max {
				r, l = end, start
				max = end - start
			}
		}
		if i+2 < slen && s[i] == s[i+2] {
			start, end := substr(s, i, i+2)
			if end-start > max {
				r, l = end, start
				max = end - start
			}
		}
	}
	return s[l : r+1]
}

func substr(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
