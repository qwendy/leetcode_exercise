package leetcode

// Given a string, find the length of the longest substring without repeating characters.

// Examples:

// Given "abcabcbb", the answer is "abc", which the length is 3.

// Given "bbbbb", the answer is "b", with the length of 1.

// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
func lengthOfLongestSubstring(s string) int {
	hash := make(map[byte]int)
	max := 0
	left := -1
	for i := 0; i < len(s); i++ {
		if index, ok := hash[s[i]]; ok {
			if index >= left {
				left = index
			}
		}
		hash[s[i]] = i
		if i-left > max {
			max = i - left
		}
	}
	return max
}
