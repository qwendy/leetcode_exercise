package practice

func shellsort(array []int) {
	n := len(array)
	h := 1
	// 1, 4, 13
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < len(array); i = i + h {
			for j := i - h; j >= 0 && array[j] > array[j+h]; j = j - h {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
		h = h / 3
	}
}
