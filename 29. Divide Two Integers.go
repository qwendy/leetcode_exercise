// https://leetcode.com/problems/divide-two-integers/description/
package leetcode

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 || (dividend == -2147483648 && divisor == -1) {
		return math.MaxInt32
	}
	var cur = uint(0)
	var sign = 1
	var divisorSign = 1
	var dividendSign = 1
	if divisor < 0 {
		divisorSign = -1
	}
	if dividend < 0 {
		dividendSign = -1
	}
	if divisorSign != dividendSign {
		sign = -1
	}
	if divisorSign == -1 {
		divisor = -divisor
	}
	if dividendSign == -1 {
		dividend = -dividend
	}

	sum := 0
	for {
		cur = 0
		for divisor<<(cur+1) <= dividend {
			cur++
		}
		for dividend-(divisor<<cur) >= 0 {
			dividend -= divisor << cur
			sum += 1 << cur
		}
		if cur == 0 {
			break
		}
	}
	if sign == -1 {
		return -sum
	}
	return sum
}
