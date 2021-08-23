package practise

import "sort"

//分发饼干
////Assume you are an awesome parent and want to give your children some cookies.
//But, you should give each child at most one cookie.
//
// Each child i has a greed factor g[i], which is the minimum size of a cookie t
//hat the child will be content with; and each cookie j has a size s[j]. If s[j] >
//= g[i], we can assign the cookie j to the child i, and the child i will be conte
//nt. Your goal is to maximize the number of your content children and output the
//maximum number.
//
//
// Example 1:
//
//
//Input: g = [1,2,3], s = [1,1]
//Output: 1
//Explanation: You have 3 children and 2 cookies. The greed factors of 3 childre
//n are 1, 2, 3.
//And even though you have 2 cookies, since their size is both 1, you could only
// make the child whose greed factor is 1 content.
//You need to output 1.
//
//
// Example 2:
//
//
//Input: g = [1,2], s = [1,2,3]
//Output: 2
//Explanation: You have 2 children and 3 cookies. The greed factors of 2 childre
//n are 1, 2.
//You have 3 cookies and their sizes are big enough to gratify all of the childr
//en,
//You need to output 2.
//
//
//
// Constraints:
//
//
// 1 <= g.length <= 3 * 104
// 0 <= s.length <= 3 * 104
// 1 <= g[i], s[j] <= 231 - 1
//
// Related Topics Array Greedy Sorting
// 👍 997 👎 136

//leetcode submit region begin(Prohibit modification and deletion)
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child, cookie := 0, 0
	for child < len(g) && cookie < len(s) {
		if s[cookie] >= g[child] {
			child++
		}
		cookie++
	}
	return child
}

//leetcode submit region end(Prohibit modification and deletion)
