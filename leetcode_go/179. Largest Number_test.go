package practice

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: []int{3, 30, 34, 5, 9},
			},
			want: "9534330",
		},
		{
			args: args{
				nums: []int{20, 1},
			},
			want: "201",
		},
		{
			args: args{
				nums: []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247},
			},
			want: "9609938824824769735703560743981399",
		},
		{
			args: args{
				nums: []int{121, 12},
			},
			want: "12121",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
