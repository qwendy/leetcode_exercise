/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
package practice

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			return -1
		}
		ok := true
		for k := 0; k+i < len(haystack) && k < len(needle); k++ {
			if haystack[k+i] != needle[k] {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	return -1
}

// @lc code=end
