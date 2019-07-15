package practice

import "testing"

func Test_shellsort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				array: []int{3, 1, 4, 1, 5, 9, 2, 6},
			},
		},
		{
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shellSort(tt.args.array)
			if !IsSorted(tt.args.array, 1) {
				t.Error(tt.args.array)
			}
		})
	}
}
