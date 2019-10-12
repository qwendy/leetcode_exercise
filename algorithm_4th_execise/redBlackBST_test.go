package practice

import (
	"fmt"
	"testing"
)

func Test_redBlackBST_putNode(t *testing.T) {
	rb := &redBlackBST{}
	for i, v := range []int{1, 2, 3, 4, -1, -2, -3, -4} {
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
