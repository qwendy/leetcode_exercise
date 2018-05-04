// https://leetcode-cn.com/explore/byteerview/card/top-byteerview-questions-easy/1/array/30/

package array

import "testing"

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				board: [][]byte{[]byte{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
					[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
