/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (37.22%)
 * Likes:    1374
 * Dislikes: 3190
 * Total Accepted:    437.8K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
 * one sorted array.
 *
 * Note:
 *
 *
 * The number of elements initialized in nums1 and nums2 are m and n
 * respectively.
 * You may assume that nums1 has enough space (size that is greater or equal to
 * m + n) to hold additional elements from nums2.
 *
 *
 * Example:
 *
 *
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * Output: [1,2,2,3,5,6]
 *
 */
package practice

// @lc code=start
func merge_1(nums1 []int, m int, nums2 []int, n int) {
	i, j, current := m-1, n-1, len(nums1)-1
	for current >= 0 {
		if j < 0 || (i >= 0 && nums1[i] > nums2[j]) {
			nums1[current] = nums1[i]
			i--
			current--
			continue
		}
		if i < 0 || (j >= 0 && nums1[i] <= nums2[j]) {
			nums1[current] = nums2[j]
			j--
			current--
		}
	}
}

// @lc code=end
