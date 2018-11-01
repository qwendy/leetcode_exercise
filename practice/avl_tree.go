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
	if r == -1 {
		x.left = avl.putNode(key, val, x.left)
	} else if r == 1 {
		x.right = avl.putNode(key, val, x.right)
	} else {
		x.val = val
	}
	x = avl.makeBalance(key, x)
	x.H = avl.max(avl.height(x.left), avl.height(x.right)) + 1
	return x
}

func (avl *avlTree) makeBalance(key interface{}, x *node) *node {

	balance := avl.height(x.left) - avl.height(x.right)

	if balance > 1 && Compare(key, x.left.key) < -1 {
		avl.leftRotate(x)
	}
	if balance > 1 && Compare(key, x.left.key) > 1 {
		avl.rightRotate(x.left)
		avl.leftRotate(x)
	}
	if balance < -1 && Compare(key, x.right.key) < -1 {
		avl.leftRotate(x.right)
		avl.rightRotate(x)
	}
	if balance < -1 && Compare(key, x.right.key) < -1 {
		avl.rightRotate(x)
	}
	return x
}

func (avl *avlTree) leftRotate(y *node) *node {
	x := y.left
	t := x.right
	x.right = y
	y.left = t

	y.H = avl.max(avl.height(y.left), avl.height(y.right)) + 1
	x.H = avl.max(avl.height(x.left), avl.height(x.right)) + 1
	return x
}

func (avl *avlTree) rightRotate(y *node) *node {
	x := y.right
	t := x.left
	y = x.left
	y.right = t

	y.H = avl.max(avl.height(y.left), avl.height(y.right)) + 1
	x.H = avl.max(avl.height(x.left), avl.height(x.right)) + 1
	return x
}

func (avl *avlTree) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
