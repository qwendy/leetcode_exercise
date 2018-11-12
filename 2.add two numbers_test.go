package practice

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n3 := &ListNode{
		Val: 3,
	}
	n1.Next = n2
	n2.Next = n3

	type args struct {
		l *ListNode
	}
	tests := []struct {
		name     string
		args     args
		wantHead *ListNode
	}{
		// TODO: Add test cases.
		{
			args: args{
				l: n1,
			},
			wantHead: &ListNode{
				Val:  3,
				Next: n2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHead := reverseList(tt.args.l); !reflect.DeepEqual(gotHead, tt.wantHead) {
				t.Errorf("reverseList() = %v, want %v", gotHead, tt.wantHead)
			}
		})
	}
}

func Test_addTwoNumbers(t *testing.T) {
	l1 := makeList([]int{2, 4, 3})
	l2 := makeList([]int{5, 6, 4})
	s3 := []int{7, 0, 8}
	l3 := makeList(s3)
	l4 := makeList([]int{1, 8})
	l5 := makeList([]int{0})
	l6 := makeList([]int{1, 8})
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
			want: l3,
		},
		{
			args: args{
				l1: l4,
				l2: l5,
			},
			want: l6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			for got != nil {
				if got.Val != tt.want.Val {
					t.Errorf("error")
					return
				}
				got = got.Next
				tt.want = tt.want.Next
			}
		})
	}
}
