package practice

import "testing"

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
