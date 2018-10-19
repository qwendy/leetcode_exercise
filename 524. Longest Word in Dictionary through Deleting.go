// https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/description/
package leetcode

func findLongestWord(s string, d []string) string {

	var r string
	for _, ss := range d {
		ok := InGivenString(ss, s)
		if ok {
			if len(ss) > len(r) || (len(ss) == len(r) && ss < r) {
				r = ss
			}
		}
	}
	return r
}

func InGivenString(word, given string) bool {
	i := 0
	for j := 0; j < len(word); j++ {
		ok := false
		for ; i < len(given); i++ {
			if given[i] == word[j] {
				i++
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
