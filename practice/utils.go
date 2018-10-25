package practice

import "reflect"

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

type ComparableValue interface {
	Compare(ComparableValue) bool
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
