package practice

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
				nums:  []int{1, 5, 1, 1, 6, 4},
				k:     5,
				left:  0,
				right: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: []int{1, 5, 1, 1, 6, 4},
			},
		},
		{
			args: args{
				nums: []int{1, 3, 2, 2, 2, 1, 1, 3, 1, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
			if !checkWiggle(tt.args.nums) {
				t.Error(tt.args.nums)
			}
		})
	}
}
