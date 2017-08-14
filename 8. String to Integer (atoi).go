package leetcode

// Implement atoi to convert a string to an integer.

// Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

// Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

// Update (2015-02-10):
// The signature of the C++ function had been updated. If you still see your function signature accepts a const char * argument, please click the reload button  to reset your code definition.
func myAtoi(str string) int {
	result := int64(0)
	times := int64(1)
	s := 0
	out := false

	for index := len(str) - 1; index >= 0; index-- {
		if str[index] < '0' || str[index] > '9' {
			if str[index] == '-' {
				s++
				result = -result
			}
			if str[index] == '+' {
				s++
			}
			if s == 2 {
				return 0
			}
			if s == 1 {
				continue
			}
			if str[index] == ' ' && (index == 0 || str[index-1] == ' ') {
				continue
			}
			times = 1
			result = 0
			continue
		}
		if !out {
			s := int64(str[index] - '0')
			result = result + times*(s)
			times = times * 10
		}
		if result > 2147483647 || result < -2147483648 {
			out = true
		} else {
			out = false
		}
	}
	if result > 2147483647 || result < -2147483648 {
		out = true
	} else {
		out = false
	}
	if out {
		if result > 0 {
			result = 2147483647
		} else {
			result = -2147483648
		}
	}

	return int(result)
}
