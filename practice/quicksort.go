package practice

func quicksort(nums []int, low, high int) {
	if low >= high {
		return
	}
	p := quickPartition(nums, low, high)
	quicksort(nums, low, p-1)
	quicksort(nums, p+1, high)

}

func quickPartition(nums []int, low, high int) int {
	i, j := low, high
	v := nums[i]
	for true {
		for nums[i] < v && i < high {
			i++
		}
		for nums[j] > v && j > low {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j], nums[low] = nums[low], nums[j]
	return j
}
