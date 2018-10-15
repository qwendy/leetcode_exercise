// https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/description/
package leetcode

func findLongestWord(s string, d []string) string {
	var letterCount [255]int
	for _, v := range s {
		letterCount[int(v)]++
	}
	var result string
	var r string
	for _, ss := range d {
		var lc [255]int
		ok := true
		sss := make([]rune, len(ss))
		i := 0
		for _, v := range ss {
			lc[int(v)]++
			if lc[int(v)] > letterCount[int(v)] {
				ok = false
				break
			} else {
				sss[i] = v
				i++
			}
		}
		if ok {
			ssss := string(sss)
			if len(sss) > len(result) || (len(sss) == len(result) && ssss < result) {
				result = ssss
				r = ss
			}
		}
	}
	return r
}
