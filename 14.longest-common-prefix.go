package practice

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	k := 0
	for ; k < len(strs[0]); k++ {
		stop := false
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= k || strs[i][k] != strs[0][k] {
				stop = true
				break
			}
		}
		if stop {
			break
		}
	}
	return strs[0][:k]
}

// @lc code=end
