// https://leetcode.com/problems/sort-list/description/
package leetcode

import "fmt"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l2 := split(head)

	return mergeList(sortList(head), sortList(l2))
}

func split(head *ListNode) *ListNode {
	slow, fast := head, head
	var tail *ListNode
	for fast != nil && fast.Next != nil {
		tail = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	tail.Next = nil
	return slow
}

func mergeList(left, right *ListNode) *ListNode {
	var head, cur, pre *ListNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur = left
			left = left.Next
		} else {
			cur = right
			right = right.Next
		}
		if head == nil {
			head = cur
		} else {
			pre.Next = cur
		}
		pre = cur
	}
	if left == nil {
		pre.Next = right
	} else {
		pre.Next = left
	}
	return head
}

func isListSorted(head *ListNode) bool {
	for ; head != nil && head.Next != nil; head = head.Next {
		if head.Val > head.Next.Val {
			return false
		}
	}
	return true
}

func printList(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
