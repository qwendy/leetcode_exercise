package practice

import "testing"

func Test_kthSmallest(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	kthSmallest(root, 2)

}
