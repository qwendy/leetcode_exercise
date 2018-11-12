package practice

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				x: 321,
			},
			want: 123,
		},
		{
			args: args{
				x: 1143,
			},
			want: 3411,
		},
		{
			args: args{
				x: 10,
			},
			want: 1,
		},
		{
			args: args{
				x: -10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
