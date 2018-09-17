// https://leetcode.com/problems/valid-anagram/description/
package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var table [26]int
	bs, bt := []byte(s), []byte(t)
	for i := 0; i < len(s); i++ {
		table[bs[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		table[bt[i]-'a']--
		if table[bt[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
