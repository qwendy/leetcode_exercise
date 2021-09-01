package practise

func minPatches(nums []int, n int) int {
	need, cover := 0, 0
	for i := 0; cover < n; {
		if i < len(nums) && nums[i]-1 <= cover {
			cover = cover + nums[i]
			i++
		} else {
			need++
			cover = 2*cover + 1
		}
	}
	return need
}
