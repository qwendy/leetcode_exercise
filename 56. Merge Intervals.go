// https://leetcode.com/problems/merge-intervals/description/
package practice

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	result := []Interval{}
	cur := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if cur.End >= intervals[i].Start {
			if cur.End < intervals[i].End {
				cur.End = intervals[i].End
			}
		} else {
			result = append(result, cur)
			cur = intervals[i]
		}
		if i == len(intervals)-1 {
			result = append(result, cur)
		}
	}
	return result
}
