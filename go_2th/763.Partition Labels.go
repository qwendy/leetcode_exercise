package practise

//Partition Labels
////You are given a string s. We want to partition the string into as many parts a
//s possible so that each letter appears in at most one part.
//
// Return a list of integers representing the size of these parts.
//
//
// Example 1:
//
//
//Input: s = "ababcbacadefegdehijhklij"
//Output: [9,7,8]
//Explanation:
//The partition is "ababcbaca", "defegde", "hijhklij".
//This is a partition so that each letter appears in at most one part.
//A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it split
//s s into less parts.
//
//
// Example 2:
//
//
//Input: s = "eccbbbbdec"
//Output: [10]
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 500
// s consists of lowercase English letters.
//
// Related Topics Hash Table Two Pointers String Greedy
// 👍 5241 👎 214

//leetcode submit region begin(Prohibit modification and deletion)
func partitionLabels(s string) []int {
	if len(s) == 0 {
		return []int{}
	}
	if len(s) == 1 {
		return []int{1}
	}
	unique := [][]int{{0, 0}}
	uLen := 0
	var exist [26]int
	for i, v := range s {
		if exist[int(v)-'a'] != 0 {
			for j := uLen - 1; j >= 0; j-- {
				if unique[j][0] <= exist[int(v)-'a']-1 {
					unique[j][1] = i
					uLen = j + 1
					break
				}
			}

		} else {
			uLen++
			if uLen > len(unique) {
				unique = append(unique, []int{i, i})
			} else {
				unique[uLen-1] = []int{i, i}
			}
		}
		exist[int(v)-'a'] = i + 1
	}
	count := make([]int, uLen)
	for i, v := range unique[:uLen] {
		count[i] = v[1] - v[0] + 1
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
