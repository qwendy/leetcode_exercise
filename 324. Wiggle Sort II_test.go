package leetcode

import (
	"testing"
)

func Test_findKthSmallest(t *testing.T) {
	type args struct {
		nums  []int
		k     int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums:  []int{5, 3, 2, 1, 4, 7, 8, 3},
				k:     5,
				left:  0,
				right: 7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthSmallest(tt.args.nums, tt.args.k, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("findKthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
