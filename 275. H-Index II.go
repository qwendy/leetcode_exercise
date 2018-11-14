package practice

func hIndex_275(citations []int) int {
	start, end := 0, len(citations)-1
	r := 0
	for start <= end {
		mid := (start + end) / 2
		if citations[mid] >= len(citations)-mid {
			end = mid - 1
			r = len(citations) - mid
		} else {
			start = mid + 1
		}
	}
	return r
}
