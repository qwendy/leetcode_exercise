package practice

import (
	"testing"
)

func Test_mergesort(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: []int{2, 3, 4, 1, 3},
				low:  0,
				high: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergesort(tt.args.nums, tt.args.low, tt.args.high)
			if !IsSorted(tt.args.nums, 1) {
				t.Error(tt.args.nums)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
		mid  int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: []int{4, 5, 6, 1, 2, 3},
				low:  0,
				high: 5,
				mid:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums, tt.args.low, tt.args.high, tt.args.mid)
			if !IsSorted(tt.args.nums, 1) {
				t.Error(tt.args.nums)
			}
		})
	}
}
