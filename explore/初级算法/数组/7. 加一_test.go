// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/27/

package array

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				digits: []int{9, 9, 9},
			},
			want: []int{1, 0, 0, 0},
		},
		{
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
