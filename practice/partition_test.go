package practice

import "testing"

func TestTreeWayPartition(t *testing.T) {
	type args struct {
		nums []int
		mid  int
	}
	tests := []struct {
		name      string
		args      args
		wantLeft  int
		wantRight int
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: []int{5, 64, 3, 2, 1, 3, 9, 8, 5, 1, 2},
				mid:  4,
			},
			wantLeft:  5,
			wantRight: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft, gotRight := TreeWayPartition(tt.args.nums, tt.args.mid)
			if gotLeft != tt.wantLeft {
				t.Errorf("TreeWayPartition() gotLeft = %v, want %v, nums: %v", gotLeft, tt.wantLeft, tt.args.nums)
			}
			if gotRight != tt.wantRight {
				t.Errorf("TreeWayPartition() gotRight = %v, want %v, nums: %v", gotRight, tt.wantRight, tt.args.nums)
			}
		})
	}
}
