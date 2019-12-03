/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
<<<<<<< HEAD
 *
 * https://leetcode-cn.com/problems/reverse-bits/description/
 *
 * algorithms
 * Easy (51.52%)
 * Likes:    104
 * Dislikes: 0
 * Total Accepted:    24.6K
 * Total Submissions: 47.7K
 * Testcase Example:  '00000010100101000001111010011100'
 *
 * 颠倒给定的 32 位无符号整数的二进制位。
 *
 *
 *
 * 示例 1：
 *
 * 输入: 00000010100101000001111010011100
 * 输出: 00111001011110000010100101000000
 * 解释: 输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
 * ⁠     因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。
 *
 * 示例 2：
 *
 * 输入：11111111111111111111111111111101
 * 输出：10111111111111111111111111111111
 * 解释：输入的二进制串 11111111111111111111111111111101 表示无符号整数 4294967293，
 * 因此返回 3221225471 其二进制表示形式为 10101111110010110010011101101001。
 *
 *
 *
 * 提示：
 *
 *
 * 请注意，在某些语言（如
 * Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
 * 在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数
 * -1073741825。
 *
 *
 *
 *
 * 进阶:
 * 如果多次调用这个函数，你将如何优化你的算法？
 *
 */
package practice
// @lc code=start
func reverseBits(num uint32) uint32 {
	var sum uint32
	for num != 0 {
		sum += num % 2
		sum = sum << 1
		num  = num / 2
	}
	return sum
}
// @lc code=end

=======
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
>>>>>>> afe370cb07d2388fc9742f35a90288777a98595d
