package practice

import (
	"fmt"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	a := []int{3, 1, 2, 4, 1, 5, 9, 2, 6}
	for _, v := range a {
		pq.insert(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(pq.delMax())
	}
	t.Error("xxx")
}
