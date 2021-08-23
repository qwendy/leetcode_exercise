// https://leetcode.com/problems/h-index/description/
package practice

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	for i, v := range citations {
		if i+1 > v {
			return i
		}
	}
	return len(citations)
}
