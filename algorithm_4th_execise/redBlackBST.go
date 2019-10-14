package practice

const (
	RED   = true
	BLACK = false
)

type redBlackBST struct {
	root *node
}

func (rb *redBlackBST) isRed(x *node) bool {
	if x == nil {
		return false
	}
	return x.color
}

func (rb *redBlackBST) get(key interface{}, x *node) interface{} {
	if x == nil {
		return nil
	}
	return x.val
}

func (rb *redBlackBST) getNode(key interface{}, x *node) *node {
	if x == nil {
		return x
	}
	r := Compare(key, x.key)
	if r < 0 {
		return rb.getNode(key, x.left)
	} else if r > 0 {
		return rb.getNode(key, x.right)
	} else {
		return x
	}
}

func (rb *redBlackBST) height(x *node) int {
	if x == nil {
		return 0
	}
	return x.H
}
func (rb *redBlackBST) size(x *node) int {
	if x == nil {
		return 0
	}
	return x.N
}
func (rb *redBlackBST) put(key, val interface{}) {
	rb.root = rb.putNode(key, val, rb.root)
	rb.root.color = RED
}
func (rb *redBlackBST) putNode(key, val interface{}, x *node) *node {
	if x == nil {
		return &node{key: key, val: key, H: 1, N: 1, color: RED}
	}
	r := Compare(key, x.key)
	if r < 0 {
		x.left = rb.putNode(key, val, x.left)
	} else if r > 0 {
		x.right = rb.putNode(key, val, x.right)
	} else {
		x.val = val
	}

	x = rb.balance(x)

	rb.setHeight(x)
	rb.setSize(x)
	return x
}

func (rb *redBlackBST) balance(x *node) *node {
	if rb.isRed(x.right) && !rb.isRed(x.left) {
		x = rb.leftRotate(x)
	}
	if rb.isRed(x.left) && rb.isRed(x.left.left) {
		x = rb.rightRotate(x)
	}
	// 双红， 则翻转颜色
	if rb.isRed(x.left) && rb.isRed(x.right) {
		x = rb.flipColor(x)
	}
	return x
}

func (rb *redBlackBST) rightRotate(y *node) *node {
	x := y.left
	t := x.right
	x.right = y
	y.left = t

	c := y.color
	y.color = x.color
	x.color = c

	rb.setHeight(y)
	rb.setHeight(x)

	rb.setSize(y)
	rb.setSize(x)
	return x
}

func (rb *redBlackBST) leftRotate(x *node) *node {
	y := x.right
	y.left = x
	y.N, x.N = rb.height(x), rb.height(y)
	x.color = RED
	x.right = nil
	y.color = BLACK
	return y
}

func (rb *redBlackBST) flipColor(x *node) *node {
	if x == nil {
		return x
	}
	x.color = !x.color
	if x.left != nil {
		x.left.color = !x.left.color
	}
	if x.right != nil {
		x.right.color = !x.right.color
	}
	return x
}

func (rb *redBlackBST) setHeight(x *node) {
	if x == nil {
		return
	}
	max := rb.height(x.left)
	if max < rb.height(x.right) {
		max = rb.height(x.right)
	}
	x.H = max + 1
}

func (rb *redBlackBST) setSize(x *node) {
	if x == nil {
		return
	}
	leftSize := rb.size(x.left)
	rightSize := rb.size(x.right)

	x.N = leftSize + rightSize + 1
}

func (rb *redBlackBST) delete(key interface{}) {
	if rb.root == nil {
		return
	}
	// 仿照不在root的时候的2-node情况，归一化处理
	if !rb.isRed(rb.root.left) && !rb.isRed(rb.root.right) {
		rb.root.color = RED
	}
	rb.root = rb.deleteNode(key, rb.root)
	if rb.root != nil {
		rb.root.color = BLACK
	}
}

func (rb *redBlackBST) deleteNode(key interface{}, x *node) *node {
	if x == nil {
		return nil
	}
	r := Compare(key, x.key)
	if r < 0 {
		if x.left == nil {
			return nil
		}
		// 当left node为RED时， 下一个节点仍处于3-node中，不需要变化。反之，确定下一个节点为2-node则需要变化
		if !rb.isRed(x.left) && !rb.isRed(x.left.left) {
			rb.moveRedLeft(x)
		}
		x.left = rb.deleteNode(key, x.left)
	} else {
		// 旋转，不论是删除本节点还是下一个（右）节点，都确定处于3-node中
		if rb.isRed(x.left) {
			rb.rightRotate(x)
		}
		if r == 0 && x.right == nil {
			return nil
		}
		if r == 0 {
			minNode := rb.min(x.right)
			x.key, x.val = minNode.key, minNode.val
			x.right = rb.deleteNode(minNode.key, x.right)
			return x
		}
		if x.right == nil {
			return nil
		}
		if !rb.isRed(x.right) && !rb.isRed(x.right.left) {
			rb.moveRedRight(x)
		}
		x.right = rb.deleteNode(key, x.right)
	}
	rb.balance(x)
	return x
}

func (rb *redBlackBST) moveRedLeft(x *node) *node {
	if x == nil {
		return nil
	}
	rb.flipColor(x)
	if rb.isRed(x.right.left) {
		x.right = rb.rightRotate(x.right)
		x = rb.leftRotate(x)
	}
	return x
}
func (rb *redBlackBST) moveRedRight(x *node) *node {
	if x == nil {
		return nil
	}
	rb.flipColor(x)
	if rb.isRed(x.left.left) {
		x = rb.rightRotate(x)
	}
	return x
}

func (rb *redBlackBST) min(x *node) *node {
	if x == nil {
		return nil
	}
	for x.left != nil {
		x = x.left
	}
	return x
}
