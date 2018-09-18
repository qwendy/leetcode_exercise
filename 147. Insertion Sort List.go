// https://leetcode.com/problems/insertion-sort-list/description/

package leetcode

func insertionSortList(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}

	cur := head
	for cur != nil && cur.Next != nil {
		p := cur.Next
		if cur.Val <= p.Val {
			cur = p
			continue
		}
		cur.Next = p.Next

		pre := headPre
		next := headPre.Next

		for next.Val < p.Val {
			pre = next
			next = next.Next
		}
		pre.Next = p
		p.Next = next
	}
	return headPre.Next
}

func insertionSortListBad(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := head
	p := head.Next
	for p != nil {
		if head.Val > p.Val {
			last.Next = p.Next
			p.Next = head
			head = p
			p = last.Next
			continue
		}
		ok := false
		for iter := head; iter.Next != nil && iter.Next != p; iter = iter.Next {
			if iter.Next.Val > p.Val {
				last.Next = p.Next
				p.Next = iter.Next
				iter.Next = p
				p = last.Next
				ok = true
				break
			}
		}
		if !ok {
			last = p
			p = p.Next
		}

	}
	return head
}
