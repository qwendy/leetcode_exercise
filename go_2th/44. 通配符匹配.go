package practise

func isMatch(s string, p string) bool {
	for len(s) > 0 && len(p) > 0 && p[len(p)-1] != '*' {
		if match(s[len(s)-1], p[len(p)-1]) {
			p = p[:len(p)-1]
			s = s[:len(s)-1]
		} else {
			return false
		}
	}
	if len(p) == 0 {
		return len(s) == 0
	}
	sIndex, pIndex := 0, 0
	sMark, pMark := -1, -1
	for sIndex < len(s) && pIndex < len(p) && pMark < len(p) {
		if p[pIndex] == '*' {
			pIndex++
			sMark, pMark = sIndex, pIndex
		} else if match(s[sIndex], p[pIndex]) {
			sIndex++
			pIndex++
		} else if sMark != -1 && sMark+1 < len(s) {
			sMark++
			pIndex = pMark
			sIndex = sMark
		} else {
			return false
		}
	}
	return allStar(p[pIndex:])
}

func match(a, b byte) bool {
	if a == b || b == '?' {
		return true
	}
	return false
}
func allStar(p string) bool {
	for i, _ := range p {
		if p[i] != '*' {
			return false
		}
	}
	return true
}
