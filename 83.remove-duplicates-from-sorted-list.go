/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (43.60%)
 * Likes:    958
 * Dislikes: 87
 * Total Accepted:    375.5K
 * Total Submissions: 861.1K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package practice

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head
	next := head.Next
	for next != nil {
		for next != nil && next.Val == current.Val {
			next = next.Next
		}
		current.Next = next
		current = current.Next
	}
	return head
}

// @lc code=end
