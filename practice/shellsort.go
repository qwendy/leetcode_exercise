package practice

func insertSort(a []int, low, high int) {
	for i := low + 1; i < high && i < len(a); i++ {
		for j := i - 1; j >= 0 && a[j] > a[j+1]; j-- {
			a[j], a[j+1] = a[j+1], a[j]
		}
	}
}

func shellSort(a []int) {
	h := 1
	n := len(a)
	for h*3+1 < n {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j-h >= 0 && a[j] < a[j-h]; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h /= 3
	}
}
