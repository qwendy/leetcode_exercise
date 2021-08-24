package practise

import (
	"sort"
)

//Non-overlapping Intervals
////Given an array of intervals intervals where intervals[i] = [starti, endi], ret
//urn the minimum number of intervals you need to remove to make the rest of the i
//ntervals non-overlapping.
//
//
// Example 1:
//
//
//Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
//Output: 1
//Explanation: [1,3] can be removed and the rest of the intervals are non-overla
//pping.
//
//
// Example 2:
//
//
//Input: intervals = [[1,2],[1,2],[1,2]]
//Output: 2
//Explanation: You need to remove two [1,2] to make the rest of the intervals no
//n-overlapping.
//
//
// Example 3:
//
//
//Input: intervals = [[1,2],[2,3]]
//Output: 0
//Explanation: You don't need to remove any of the intervals since they're alrea
//dy non-overlapping.
//
//
//
// Constraints:
//
//
// 1 <= intervals.length <= 105
// intervals[i].length == 2
// -5 * 104 <= starti < endi <= 5 * 104
//
// Related Topics Array Dynamic Programming Greedy Sorting
// 👍 2525 👎 73

//leetcode submit region begin(Prohibit modification and deletion)
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 0
	last := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < last {
			count++
		} else {
			last = intervals[i][1]
		}

	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
