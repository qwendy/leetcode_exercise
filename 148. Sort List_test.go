package leetcode

import (
	"testing"
)

func Test_mergeList(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 3}
	l3 := &ListNode{Val: 2}
	l4 := &ListNode{Val: 4}
	l1.Next = l3
	l2.Next = l4

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			args: args{
				l1: l1,
				l2: l2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeList(tt.args.l1, tt.args.l2); !isListSorted(tt.args.l1) {
				printList(tt.args.l1)
				t.Errorf("mergeList() = %v, want %v", got, tt.want)
			}
		})
	}
}
