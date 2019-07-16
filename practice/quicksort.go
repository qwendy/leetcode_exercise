package practice

func quicksort(nums []int, low, high int) {
	if low+15 >= high {
		insertSort(nums, low, high)
		return
	}
	p := quickPartition(nums, low, high)
	quicksort(nums, low, p-1)
	quicksort(nums, p+1, high)

}

func quickPartition(nums []int, low, high int) int {
	i, lt, gt := low, low+1, high

}
