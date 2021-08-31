package practise

func jump(nums []int) int {
	step := 0
	most := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		m := nums[i] + i
		if m > most {
			most = m
		}
		if i == end {
			step++
			end = most
		}
	}
	return step
}
