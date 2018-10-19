package practice

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
