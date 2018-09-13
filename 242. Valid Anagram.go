package leetcode

// https://leetcode.com/problems/valid-anagram/description/
func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return true
	}
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		if _, ok := m[v]; ok {
			m[v]--
			if m[v] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
