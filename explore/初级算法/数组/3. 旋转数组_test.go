// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/23/

package array

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{-1},
				k:    2,
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			t.Errorf("sfs:%v", tt.args.nums)
		})
	}

}
