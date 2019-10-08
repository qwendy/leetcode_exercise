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

func partition1(nums []int, low, high int) int {
	i, j := low, high+1
	v := nums[low]
	for true {
		for {
			i++
			if i > high || nums[i] >= v {
				break
			}
		}
		for {
			j--
			if j < low || nums[j] <= v {
				break
			}
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[low], nums[j] = nums[j], nums[low]
	return j
}
func ThreeWayPartition(nums []int, mid int) (left, right int) {
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
func twp(nums []int, low, high int) (left, right int) {
	i, left, right := low+1, low, right+1
	v := nums[low]
	for left < right && i <= high {
		if nums[i] < v {
			left++
			nums[left], nums[i] = nums[i], nums[left]
			i++
		}
		if nums[i] == v {
			i++
		}
		if nums[i] > v {
			right--
			nums[i], nums[right] = nums[right], nums[i]
		}
	}

	return
}
