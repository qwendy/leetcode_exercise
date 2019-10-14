package practice

import (
	"fmt"
	"testing"
)

func Test_redBlackBST_putNode(t *testing.T) {
	rb := &redBlackBST{}
	for i, v := range []int{1, 2, 3, 4, -1, -2, -3, -4} {
		rb.put(v, i)

	}
	array := make([]interface{}, 0)
	array = InOrder(array, rb.root)
	fmt.Println(array)

	for i, v := range []int{1, 2, 3, 4, -1, -2, -3, -4} {
		rb.delete(v)
		array = make([]interface{}, 0)
		array = InOrder(array, rb.root)
		fmt.Println(i, v, array)

	}

}
