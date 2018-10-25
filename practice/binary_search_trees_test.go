package practice

import (
	"fmt"
	"testing"
)

func Test_binarySearchTree_put(t *testing.T) {
	bst := binarySearchTree{}
	for i, v := range []int{9, 8, 3, 11, 14, 2, 10, 27} {
		bst.put(v, i)
	}
	bst.printKeys(bst.root)

	bst.root = bst.deleteMin(bst.root)
	fmt.Println("-----")
	bst.printKeys(bst.root)

	bst.root = bst.deleteMax(bst.root)
	fmt.Println("-----")
	bst.printKeys(bst.root)

	bst.delete(9)
	fmt.Println("------")
	bst.printKeys(bst.root)

	fmt.Println("isBST:", bst.isBST(bst.root))

	fmt.Println("height:", bst.height(bst.root))
}
