// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/27/

package array

func plusOne(digits []int) []int {
	result := make([]int, len(digits))
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] >= 10 {
			carry = digits[i] / 10
			digits[i] = digits[i] % 10
		} else {
			carry = 0
		}
		result[i] = digits[i]
	}
	if carry > 0 {
		result = append([]int{1}, result...)
	}
	return result
}
