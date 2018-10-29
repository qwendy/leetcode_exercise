// https://www.geeksforgeeks.org/avl-tree-set-1-insertion/
package practice

type avlTree struct {
	root *node
}

func (avl *avlTree) height(x *node) int {
	if x == nil {
		return 0
	}
	return x.H
}

func (avl *avlTree) put(key, val interface{}) {
	avl.root = avl.putNode(key, val, avl.root)
}

func (avl *avlTree) putNode(key, val interface{}, x *node) *node {
	if x == nil {
		return &node{key: key, val: val, H: 1}
	}

	r := Compare(key, x.key)
	if r < -1 {
		avl.putNode(key, val, x)
	} else if r > 1 {
		avl.putNode(key, val, x)
	} else {
		x.val = val
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	x.H = max(avl.height(x.left), avl.height(x.right)) + 1

	balance := avl.height(x.left) - avl.height(x.right)

	// TODO:  left rotate
	if balance > 1 && Compare(key, x.left.key) < -1 {
	}

	return x
}
