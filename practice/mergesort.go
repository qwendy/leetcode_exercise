package practice

func mergesort(nums []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergesort(nums, low, mid)
	mergesort(nums, mid+1, high)
	merge(nums, low, high, mid)

}

func merge(nums []int, low, high, mid int) {
	aux := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		aux[i] = nums[i]
	}
	for k, i, j := low, low, mid+1; k <= high; k++ {
		if i > mid {
			nums[k] = aux[j]
			j++
		} else if j > high {
			nums[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			nums[k] = aux[i]
			i++
		} else {
			nums[k] = aux[j]
			j++
		}
	}
}
