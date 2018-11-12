package practice

import (
	"testing"
)

func Test_checkPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				s: "123211",
			},
			want: false,
		},
		{
			args: args{
				s: "12321",
			},
			want: true,
		},
		{
			args: args{
				s: "1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPalindrome(tt.args.s); got != tt.want {
				t.Errorf("checkPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				s: "123211",
			},
			want: "12321",
		},
		{
			args: args{
				s: "12321",
			},
			want: "12321",
		},
		{
			args: args{
				s: "1",
			},
			want: "1",
		},
		{
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
