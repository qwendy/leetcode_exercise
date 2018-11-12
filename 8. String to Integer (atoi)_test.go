package practice

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				str: "123",
			},
			want: 123,
		},
		{
			args: args{
				str: "-123",
			},
			want: -123,
		},
		{
			args: args{
				str: "+-2",
			},
			want: 0,
		},
		{
			args: args{
				str: "     +004500",
			},
			want: 4500,
		},
		{
			args: args{
				str: "  -0012a42",
			},
			want: -12,
		},
		{
			args: args{
				str: "123  456",
			},
			want: 123,
		},
		{
			args: args{
				str: "   010",
			},
			want: 10,
		},
		{
			args: args{
				str: "18446744073709551617",
			},
			want: 2147483647,
		},
		{
			args: args{
				str: "  +b12102370352",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
