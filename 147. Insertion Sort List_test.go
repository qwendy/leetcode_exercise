// https://leetcode.com/problems/insertion-sort-list/description/

package practice

import (
	"fmt"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	l1 := ListNode{
		Val: 4,
	}
	l2 := ListNode{
		Val: 2,
	}
	l3 := ListNode{
		Val: 1,
	}
	l4 := ListNode{
		Val: 3,
	}
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			args: args{
				head: &l1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertionSortList(tt.args.head)
			for ; got != nil; got = got.Next {
				fmt.Println(got.Val)
			}
			t.Error("stop")
		})
	}
}
