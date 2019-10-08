package practice

import (
	"fmt"
	"testing"
)

func Test_redBlackBST_putNode(t *testing.T) {
	rb := &redBlackBST{}
	for i, v := range []int{3, 1, 4, 1, 5, 9, 2, 6} {
		rb.put(v, i)

		array := make([]interface{}, 0)
		array = InOrder(array, rb.root)

		prearray := make([]interface{}, 0)
		prearray = PreOrder(prearray, rb.root)
		fmt.Println(array, prearray)
	}

	rb.delete(3)

	array := make([]interface{}, 0)
	array = InOrder(array, rb.root)
	fmt.Println(array)
}
