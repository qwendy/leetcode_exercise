package practice

import "fmt"

type node struct {
	key   interface{}
	val   interface{}
	left  *node
	right *node

	H int
	N int
}

type binarySearchTree struct {
	root *node
}

func (bst *binarySearchTree) size(x *node) int {
	if x == nil {
		return 0
	} else {
		return x.N
	}
}

// use recursion to calculate the size
func (bst *binarySearchTree) sizeV2(x *node) int {
	if x == nil {
		return 0
	}
	return bst.sizeV2(x.left) + bst.sizeV2(x.right) + 1
}

func (bst *binarySearchTree) get(key interface{}) interface{} {
	x := bst.findNode(key, bst.root)
	if x == nil {
		return nil
	} else {
		return x.val
	}
}

func (bst *binarySearchTree) put(key, val interface{}) {
	bst.root = bst.putNode(key, val, bst.root)
}

func (bst *binarySearchTree) findNode(key interface{}, x *node) *node {
	if x == nil {
		return x
	}
	r := Compare(key, x.key)
	if r == -1 {
		return bst.findNode(key, x.left)
	} else if r == 1 {
		return bst.findNode(key, x.right)
	} else {
		return x
	}
}

func (bst *binarySearchTree) putNode(key, val interface{}, x *node) *node {
	if x == nil {
		return &node{
			key: key,
			val: val,
			N:   1,
		}
	}
	r := Compare(key, x.key)
	if r == -1 {
		x.left = bst.putNode(key, val, x.left)
	} else if r == 1 {
		x.right = bst.putNode(key, val, x.right)
	} else {
		x.val = val
	}
	x.N = bst.size(x.left) + bst.size(x.right) + 1
	return x
}
func (bst *binarySearchTree) min(x *node) *node {
	if x == nil {
		return x
	}
	if x.left != nil {
		return bst.min(x.left)
	}
	return x
}
func (bst *binarySearchTree) max(x *node) *node {
	if x == nil {
		return x
	}
	if x.right != nil {
		return bst.max(x.right)
	}
	return x
}
func (bst *binarySearchTree) deleteMin(x *node) *node {
	if x == nil {
		return nil
	}
	if x.left == nil {
		return x.right
	}
	x.left = bst.deleteMin(x.left)
	x.N = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *binarySearchTree) deleteMax(x *node) *node {
	if x == nil {
		return nil
	}
	if x.right == nil {
		return x.left
	}
	x.right = bst.deleteMax(x.right)
	x.N = bst.size(x.left) + bst.size(x.right) + 1
	return x
}
func (bst *binarySearchTree) delete(key interface{}) {
	bst.root = bst.deleteKey(key, bst.root)
}

func (bst *binarySearchTree) deleteKey(key interface{}, x *node) *node {
	if x == nil {
		return x
	}
	r := Compare(key, x.key)
	if r < 0 {
		x.left = bst.deleteKey(key, x.left)
	} else if r > 0 {
		x.right = bst.deleteKey(key, x.right)
	} else {
		if x.left == nil {
			return x.right
		}
		if x.right == nil {
			return x.left
		}
		t := x
		x = bst.min(t.right)
		x.right = bst.deleteMin(t.right)
		x.left = t.left
	}
	x.N = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *binarySearchTree) printKeys(x *node) {
	if x == nil {
		return
	}
	bst.printKeys(x.left)
	fmt.Println(x.key)
	bst.printKeys(x.right)
}

func (bst *binarySearchTree) isBST(x *node) bool {
	if x == nil {
		return true
	}
	if x.left != nil {
		r := Compare(x.key, x.left.key)
		if r < 0 {
			return false
		}
		return bst.isBST(x.left)

	}
	if x.right != nil {
		r := Compare(x.key, x.right.key)
		if r > 0 {
			return false
		}
		return bst.isBST(x.right)
	}
	return true
}
