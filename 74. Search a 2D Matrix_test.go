package leetcode

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				matrix: [][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 50},
				},
				target: 13,
			},
			want: false,
		},
		{
			args: args{
				matrix: [][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 50},
				},
				target: 3,
			},
			want: true,
		},
		{
			args: args{
				matrix: [][]int{
					[]int{1},
					[]int{3},
				},
				target: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
