package practice

const (
	RED   = true
	BLACK = false
)

type redBlackBST struct {
	root *node
}

func (rb *redBlackBST) isRed(x *node) bool {
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

	x = rb.correct(x)

	rb.setHeight(x)
	rb.setSize(x)
	return x
}

func (rb *redBlackBST) correct(x *node) *node {
	if x.right != nil && rb.isRed(x.right) {
		x = rb.leftRotate(x)
	}
	if x.left != nil && x.left.left != nil && rb.isRed(x.left) && rb.isRed(x.left.left) {
		x = rb.rightRotate(x)
	}
	if x.right != nil && x.left != nil && rb.isRed(x.left) && rb.isRed(x.right) {
		x = rb.flipColor(x)
	}

	return x
}

func (rb *redBlackBST) rightRotate(y *node) *node {
	x := y.left
	t := x.right
	x.right = y
	y.left = t

	y.color = x.color
	x.color = RED

	rb.setHeight(y)
	rb.setHeight(x)

	rb.setSize(y)
	rb.setSize(x)
	return x
}

func (rb *redBlackBST) leftRotate(x *node) *node {
	y := x.right
	t := y.left
	y.left = x
	x.right = t

	x.color = y.color
	y.color = RED

	rb.setHeight(x)
	rb.setHeight(y)

	rb.setSize(x)
	rb.setSize(y)
	return y
}

func (rb *redBlackBST) flipColor(x *node) *node {
	if x.left != nil {
		x.color = BLACK
	}
	if x.right != nil {
		x.color = BLACK
	}
	x.color = RED
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
