package practice

func kthSmallest(root *TreeNode, k int) int {
	var n = 0
	_, r := PreOrder_230(root, &n, k)
	return r
}

func PreOrder_230(x *TreeNode, n *int, k int) (ok bool, r int) {
	if x == nil {
		return
	}
	ok, r = PreOrder_230(x.Left, n, k)
	if ok {
		return
	}
	*n += 1
	if *n == k {
		return true, x.Val
	}
	return PreOrder_230(x.Right, n, k)
}
