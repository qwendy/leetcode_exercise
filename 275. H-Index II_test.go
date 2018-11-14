package practice

import "testing"

func Test_hIndex_275(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				citations: []int{0, 1, 3, 5, 6},
			},
			want: 3,
		},
		{
			args: args{
				citations: []int{0, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex_275(tt.args.citations); got != tt.want {
				t.Errorf("hIndex_275() = %v, want %v", got, tt.want)
			}
		})
	}
}
