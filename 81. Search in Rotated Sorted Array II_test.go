package practice

import "testing"

func Test_search_81(t *testing.T) {
	type args struct {
		nums   []int
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
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 3,
			},
			want: false,
		},
		{
			args: args{
				nums:   []int{1, 3, 1, 1, 1},
				target: 3,
			},
			want: true,
		},
		{
			args: args{
				nums:   []int{},
				target: 3,
			},
			want: false,
		},
		{
			args: args{
				nums:   []int{1, 3},
				target: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search_81(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search_81() = %v, want %v", got, tt.want)
			}
		})
	}
}
