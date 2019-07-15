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
	pivot := nums[low]
	i, j := low, high
	for i < j {
		for ; i < j && nums[i] < pivot; i++ {
		}
		for ; j > i && nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[j], nums[low] = nums[low], nums[j]
	return j
}
