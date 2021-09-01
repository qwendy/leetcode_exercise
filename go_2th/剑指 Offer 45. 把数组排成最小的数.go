package practise

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	ns := make([]string, len(nums))
	for i, v := range nums {
		ns[i] = strconv.Itoa(v)
	}
	sort.Slice(ns, func(i, j int) bool {
		if compare(ns[i], ns[j]) == -1 {
			return true
		}
		if compare(ns[i], ns[j]) == 1 {
			return false
		}
		return false
	})
	s := ""
	for _, v := range ns {
		s += v
	}
	return s
}

func compare(a, b string) int {
	sa := a + b
	sb := b + a
	for i := 0; i < len(a)+len(b); i++ {
		if sa[i] > sb[i] {
			return 1
		}
		if sa[i] < sb[i] {
			return -1
		}
	}
	return 0
}
