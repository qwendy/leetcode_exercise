// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/28/

package array

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			t.Error(tt.args.nums)
		})
	}
}
