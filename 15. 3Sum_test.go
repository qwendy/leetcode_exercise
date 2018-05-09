// https://leetcode.com/problems/3sum/description/

package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				nums: []int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3},
			},
			want: [][]int{
				[]int{-5, 1, 4},
				[]int{-5, 2, 3},
				[]int{-3, -1, 4},
				[]int{-3, 0, 3},
				[]int{-3, 1, 2},
				[]int{2, -2, 4},
				[]int{-2, -1, 3},
				[]int{-2, 0, 2},
				[]int{-2, 1, 1},
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
