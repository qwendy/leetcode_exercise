package practice

import (
	"reflect"
	"testing"
)

func TestInOrder(t *testing.T) {
	type args struct {
		array []interface{}
		x     *node
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InOrder(tt.args.array, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
