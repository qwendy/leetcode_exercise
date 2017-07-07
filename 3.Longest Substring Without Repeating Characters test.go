package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			args: args{
				s: "abba",
			},
			want: 2,
		},
		{
			args: args{
				s: "c",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
