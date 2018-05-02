// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/26/

package array

func intersect(nums1 []int, nums2 []int) []int {
	source := nums1
	target := nums2
	result := make([]int, 0)
	if len(nums1) > len(nums2) {
		source = nums2
		target = nums1
	}
	// hash
	h := make(map[int]int, 0)
	for _, value := range source {
		h[value]++
	}
	for _, value := range target {
		if v, ok := h[value]; ok && v > 0 {
			result = append(result, value)
			h[value]--
		}
	}
	return result
}
