// 75. Sort Colors
package practice

func sortColors(nums []int) {
	left := -1
	right := len(nums)
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums) && left <= right && i < right; {
		switch {
		case nums[i] == 0:
			left++
			nums[left], nums[i] = nums[i], nums[left]
			i++
		case nums[i] == 2:
			right--
			nums[right], nums[i] = nums[i], nums[right]
		default:
			i++
		}

	}
}

// two pass
func sortColorsTwoPass(nums []int) {
	var h [3]int
	for _, v := range nums {
		h[v]++
	}
	count := 0
	for i, v := range h {
		for k := count; k < count+v; k++ {
			nums[k] = i
		}
		count += v
	}
}
