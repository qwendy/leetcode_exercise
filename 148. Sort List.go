// https://leetcode.com/problems/sort-list/description/
package leetcode

import "fmt"

func sortList(head *ListNode) *ListNode {
	headPre := &ListNode{
		Next: head,
	}
	return headPre.Next
}
func sortChild(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	headPre := &ListNode{
		Next: head,
	}
	first := head
	second := head.Next
	for first != nil && second != nil {
		first = first.Next
		second = head.Next.Next
	}
}

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	headPre := &ListNode{
		Next: l1,
	}
	pre := headPre
	next := l1
	for l2 != nil {
		t := l2
		l2 = l2.Next
		for next != nil && next.Val < t.Val {
			pre = next
			next = next.Next
		}
		pre.Next = t
		t.Next = next
	}
	return headPre.Next
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
