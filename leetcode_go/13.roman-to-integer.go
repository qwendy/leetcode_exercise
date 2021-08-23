package practice

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	var result int
	convert := func(symbol byte) int {
		var value = 0
		switch symbol {
		case 'I':
			value = 1
		case 'V':
			value = 5
		case 'X':
			value = 10
		case 'L':
			value = 50
		case 'C':
			value = 100
		case 'D':
			value = 500
		case 'M':
			value = 1000
		}
		return value
	}
	for i := 0; i < len(s); i++ {
		value := convert(s[i])
		if i+1 < len(s) && value < convert(s[i+1]) {
			value = -value
		}
		result += value
	}
	return result
}

// @lc code=end
