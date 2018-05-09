package leetcode

import "testing"

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums: []int{1, 3, 2},
			},
		},
		{
			args: args{
				nums: []int{2, 3, 1},
			},
		},
		{
			args: args{
				nums: []int{3, 2, 1},
			},
		},
		{
			args: args{
				nums: []int{1, 2, 3},
			},
		},
		{
			args: args{
				nums: []int{3, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			t.Error(tt.args.nums)
		})
	}
}
