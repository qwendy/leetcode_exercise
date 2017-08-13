package leetcode

import "strconv"

// Reverse digits of an integer.

// Example1: x = 123, return 321
// Example2: x = -123, return -321

// Have you thought about this?
// Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!

// If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.

// Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer, then the reverse of 1000000003 overflows. How should you handle such cases?

// For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

// Note:
// The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.
func reverse(x int) int {
	symbol := 1
	if x < 0 {
		x = -x
		symbol = -1
	}
	a := strconv.Itoa(x)
	b := make([]byte, len(a))
	for index := len(a) - 1; index >= 0; index-- {
		b[len(a)-1-index] = a[index]
	}
	result, err := strconv.Atoi(string(b))
	if err != nil {
		return 0
	}
	result = symbol * result
	if result > 2147483647 || result < -2147483648 {
		return 0
	}

	return result
}
