package practice

import (
	"fmt"
	"testing"
)

func Test_avlTree_put(t *testing.T) {
	avl := &avlTree{}
	for i, v := range []int{3, 1, 4, 1, 5, 9, 2, 6} {
		avl.put(v, i)
	}
	array := make([]interface{}, 0)
	array = InOrder(array, avl.root)
	fmt.Println(array)
}
