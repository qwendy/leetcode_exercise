// https://leetcode.com/problems/intersection-of-two-arrays/description/
package practice

func intersection(nums1 []int, nums2 []int) []int {
	r := make([]int, 0)
	h := make(map[int]bool, 10)
	for _, v := range nums1 {
		h[v] = true
	}
	for _, v := range nums2 {
		if h[v] {
			h[v] = false
			r = append(r, v)
		}
	}
	return r
}
