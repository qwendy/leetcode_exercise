package practise

import (
	"reflect"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	type args struct {
		people [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{people: [][]int{{8, 2}, {4, 2}, {4, 5}, {2, 0}, {7, 2}, {1, 4}, {9, 1}, {3, 1}, {9, 0}, {1, 0}}},
			want: [][]int{{1, 0}, {2, 0}, {9, 0}, {3, 1}, {1, 4}, {9, 1}, {4, 2}, {7, 2}, {8, 2}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructQueue(tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
