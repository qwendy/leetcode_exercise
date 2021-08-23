package practice

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: root,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
