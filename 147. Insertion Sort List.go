// https://leetcode.com/problems/insertion-sort-list/description/

package leetcode

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next
	for ; p != nil; p = p.Next {
		t := p.Val

		for iter := head; iter != p; iter = iter.Next {
			if t < iter.Val {
				v := iter.Val
				iter.Val = t
				iter = iter.Next

				for {
					v, iter.Val = iter.Val, v
					if iter == p {
						break
					}
					iter = iter.Next
				}
				break
			}
		}
	}
	return head
}
