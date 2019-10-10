package practice

import (
	"fmt"
	"testing"
)

func Test_binarySearchTree_put(t *testing.T) {
	bst := binarySearchTree{}
	for i, v := range []int{1, 2, 3, -1, -2, 4, 5, 6} {
		bst.put(v, i)
	}
	bst.printKeys(bst.root)

	fmt.Println("isBST:", bst.isBST(bst.root))

}
