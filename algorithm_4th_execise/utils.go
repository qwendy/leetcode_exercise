package practice

import (
	"reflect"
)

func IsSorted(array []int, order int) bool {
	for i := 0; i < len(array)-1; i++ {
		if order == 1 {
			if array[i] > array[i+1] {
				return false
			}
		} else {
			if array[i] < array[i+1] {
				return false
			}
		}

	}
	return true
}

// TODO: other type to be continued
func Compare(a, b interface{}) int {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	t := va.Type()
	switch t.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16:
		if va.Int() > vb.Int() {
			return 1
		} else if va.Int() < vb.Int() {
			return -1
		} else {
			return 0
		}
	}
	return 0
}

func PreOrder(array []interface{}, x *node) []interface{} {
	if x == nil {
		return array
	}
	array = append(array, x.key)
	array = PreOrder(array, x.left)
	array = PreOrder(array, x.right)
	return array
}

func InOrder(array []interface{}, x *node) []interface{} {
	if x == nil {
		return array
	}
	array = InOrder(array, x.left)
	array = append(array, x.key)
	array = InOrder(array, x.right)
	return array
}
