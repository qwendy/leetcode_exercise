package leetcode

// There are two sorted arrays nums1 and nums2 of size m and n respectively.

// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

// Example 1:
// nums1 = [1, 3]
// nums2 = [2]

// The median is 2.0
// Example 2:
// nums1 = [1, 2]
// nums2 = [3, 4]

// The median is (2 + 3)/2 = 2.5
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := float64(0)
	nums3 := make([]int, len(nums1)+len(nums2))
	i := 0
	j := 0
	for i < len(nums1) || j < len(nums2) {

	}
}
