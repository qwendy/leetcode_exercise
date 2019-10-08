package practice

import (
	"fmt"
	"testing"
)

func TestNewDepthFirstSearch(t *testing.T) {
	dfs := NewDepthFirstSearch()
	dfs.MakeGraph(4, [][]int{[]int{0, 1}, []int{1, 2}, []int{1, 3}, []int{2, 3}})
	path := dfs.PathTo(0, 2)
	fmt.Println(path)
}
