package practise

//去除重复字母
////给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
//
// 注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-
//distinct-characters 相同
//
//
//
// 示例 1：
//
//
//输入：s = "bcabc"
//输出："abc"
//
//
// 示例 2：
//
//
//输入：s = "cbacdcbc"
//输出："acdb"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 由小写英文字母组成
//
// Related Topics 栈 贪心 字符串 单调栈 👍 589 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicateLetters(s string) string {
	if len(s) == 1 {
		return s
	}
	exist := make([]int, 26)
	mark := make([]bool, len(s))
	for i := 0; i < len(exist); i++ {
		exist[i] = -1
	}
	for i, v := range s {
		if exist[int(v)-'a'] != -1 {
			mark[i] = true
			mark[int(v)-'a'] = true
			exist[int(v)-'a'] = i
		} else {
			exist[int(v)-'a'] = i
		}
	}
	r := ""
	for i, v := range s {
		if mark[i] {
			r += string(v)
		}
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-08-29 00:58:26
