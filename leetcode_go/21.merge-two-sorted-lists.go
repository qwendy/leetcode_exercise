/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (49.27%)
 * Likes:    2796
 * Dislikes: 409
 * Total Accepted:    722.6K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	pre := head
	current := head.Next
	for {
		if l1 == nil {
			pre.Next = l2
			return head.Next
		}
		if l2 == nil {
			pre.Next = l1
			return head.Next
		}
		if l1.Val <= l2.Val {
			current = l1
			l1 = l1.Next
		} else {
			current = l2
			l2 = l2.Next
		}
		pre.Next = current
		pre = pre.Next
	}
}

// @lc code=end
