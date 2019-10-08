package practice

import (
	"testing"
)

func Test_quicksort(t *testing.T) {
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
				nums: []int{3, 1, 4, 1, 5, 9, 2, 6},
				low:  0,
				high: 7,
			},
		},
		{
			args: args{
				nums: []int{2, 2, 2},
				low:  0,
				high: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort(tt.args.nums, tt.args.low, tt.args.high)
		})
	}
}

func Test_partition1(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: []int{3, 1, 4, 1, 5, 9, 2, 6},
				low:  0,
				high: 7,
			},

			want: 3,
		},
		{
			args: args{
				nums: []int{2, 2, 2},
				low:  0,
				high: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition1(tt.args.nums, tt.args.low, tt.args.high)
			if got != tt.want {
				t.Errorf("partition1() got = %v, want %v, nums %v", got, tt.want, tt.args.nums)
			}

		})
	}
}

func Test_twp(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name      string
		args      args
		wantLeft  int
		wantRight int
	}{
		{
			args: args{
				nums: []int{3, 1, 4, 1, 5, 9, 2, 6},
				low:  0,
				high: 7,
			},
		},
		{
			args: args{
				nums: []int{2, 2, 2},
				low:  0,
				high: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft, gotRight := twp(tt.args.nums, tt.args.low, tt.args.high)
			if gotLeft != tt.wantLeft {
				t.Errorf("twp() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if gotRight != tt.wantRight {
				t.Errorf("twp() gotRight = %v, want %v", gotRight, tt.wantRight)
			}
			t.Error(tt.args.nums)
		})
	}
}
