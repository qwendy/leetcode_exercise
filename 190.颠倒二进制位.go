/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

package practice

// @lc code=start
func reverseBits(num uint32) uint32 {
	var sum uint32
	for i := 0; i < 32; i++ {
		sum = sum << 1
		sum += num % 2
		num = num / 2
	}
	return sum
}

// @lc code=end
