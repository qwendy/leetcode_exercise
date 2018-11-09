package leetcode

import (
	"math"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				dividend: 17,
				divisor:  4,
			},
			want: 4,
		},
		{
			args: args{
				dividend: 15,
				divisor:  3,
			},
			want: 5,
		},
		{
			args: args{
				dividend: -15,
				divisor:  3,
			},
			want: -5,
		},
		{
			args: args{
				dividend: 0,
				divisor:  1,
			},
			want: 0,
		},
		{
			args: args{
				dividend: -1,
				divisor:  -1,
			},
			want: 1,
		},
		{
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			want: math.MaxInt32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
