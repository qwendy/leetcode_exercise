// https://www.geeksforgeeks.org/avl-tree-set-1-insertion/
package practice

type avlTree struct {
	root *node
}

func (avl *avlTree) height(x *node) int {
	max := avl.height(x.left)
	if max < avl.height(x.right) {
		max = avl.height(x.right)
	}
	return max + 1
}
