package practise

func validPalindrome(s string) bool {
	count := 0
	i, j := 0, len(s)-1
	for i <= j && i < len(s) && j >= 0 {
		if s[i] != s[j] {
			count++
			if count > 1 {
				break
			}
			j--
		} else {
			i++
			j--
		}
	}
	if count <= 1 {
		return true
	}
	count = 0
	for i, j := 0, len(s)-1; i <= j && i < len(s) && j >= 0; {
		if s[i] != s[j] {
			count++
			i++
			if count > 1 {
				return false
			}
		} else {
			i++
			j--
		}
	}
	return true
}
