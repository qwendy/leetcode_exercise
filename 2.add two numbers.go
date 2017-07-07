package leetcode

import "fmt"

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry = carry + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry = carry + l2.Val
			l2 = l2.Next
		}
		p.Next = NewListNode(carry%10, nil)
		carry = carry / 10
		p = p.Next
	}
	if carry != 0 {
		p.Next = NewListNode(carry, nil)
	}
	return head.Next
}
func reverseList(l *ListNode) (head *ListNode) {
	p1 := l
	p2 := l.Next
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	l.Next = nil
	return p1
}
func printListNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%v,", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func makeList(s []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, val := range s {
		p.Next = NewListNode(val, nil)
		p = p.Next
	}
	return head.Next
}
