package practice

func TreeWayPartition(nums []int, mid int) (left, right int) {
	i, j, k := 0, len(nums)-1, 0
	for i < j && k < j {
		if nums[k] < mid {
			nums[k], nums[i] = nums[i], nums[k]
			i++
			k++
		} else if nums[k] > mid {
			nums[k], nums[j] = nums[j], nums[k]
			j--
		} else {
			k++
		}
	}
	return i - 1, j
}
