package practice

func quicksort(nums []int, low, high int) {
	if low+15 >= high {
		insertSort(nums, low, high)
		return
	}
	lt, gt := quickPartition(nums, low, high)
	quicksort(nums, low, lt-1)
	quicksort(nums, gt+1, high)

}

func quickPartition(nums []int, low, high int) (int, int) {
	i, lt, gt := low+1, low, high
	v := nums[low]
	for i <= high && lt < gt {
		switch {
		case nums[i] < v:
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
			i++
		case nums[i] > v:
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		default:
			i++
		}
	}
	return lt, gt
}
