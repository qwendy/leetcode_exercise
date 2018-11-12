// https://leetcode.com/problems/largest-number/description/
package practice

import (
	"bytes"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	var ss []string
	for _, v := range nums {
		ss = append(ss, strconv.Itoa(v))
	}
	sort.Slice(ss, func(i, j int) bool {
		k := 0
		for ; k < len(ss[i]) && k < len(ss[j]); k++ {
			if ss[i][k] != ss[j][k] {
				return ss[i][k] > ss[j][k]
			}
		}
		if k < len(ss[i]) {
			for ; k < len(ss[i]); k++ {
				if ss[i][k] != ss[i][0] {
					return ss[i][k] > ss[i][0]
				}
			}
			return ss[i][1] < ss[i][0]
		}
		if k < len(ss[j]) {
			for ; k < len(ss[j]); k++ {
				if ss[j][k] != ss[j][0] {
					return !(ss[j][k] > ss[j][0])
				}
			}
			return ss[j][1] > ss[j][0]
		}
		return false
		// return ss[i]+ss[j] > ss[j]+ss[i]
	})
	var buffer bytes.Buffer
	for i, v := range ss {
		if i == 0 && v == "0" {
			return "0"
		}
		buffer.WriteString(v)
	}
	return buffer.String()
}
