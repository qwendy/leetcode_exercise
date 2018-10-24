package practice

import (
	"fmt"
	"testing"
)

func Test_symbolTable_put(t *testing.T) {
	st := &symbolTable{}
	st.put(1, 3)
	st.put(2, 1)
	fmt.Println(st.get(1))
	fmt.Println(st.get(2))
	st.delete(1)
	fmt.Println(st.get(1), st.get(2))
	t.Error("stop")
}
