package practice

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	} else if len(matrix[0]) == 0 {
		return false
	} else {
		// Find the possible rows.
		rowStart, rowMid, rowEnd := 0, 0, len(matrix)-1
		smaller, larger := 0, len(matrix[0])
		for rowStart <= rowEnd {
			rowMid = (rowStart + rowEnd) / 2
			if matrix[rowMid][0] == target {
				return true
			} else {
				if matrix[rowMid][0] < target {
					rowStart = rowMid + 1
					smaller = rowMid
				} else {
					rowEnd = rowMid - 1
				}
			}
		}

		rowStart, rowMid, rowEnd = 0, 0, len(matrix)-1
		for rowStart <= rowEnd {
			rowMid = (rowStart + rowEnd) / 2
			if matrix[rowMid][len(matrix[0])-1] == target {
				return true
			} else {
				if matrix[rowMid][len(matrix[0])-1] < target {
					rowStart = rowMid + 1
				} else {
					rowEnd = rowMid - 1
					larger = rowMid
				}
			}
		}

		for i := larger; i <= smaller; i++ {
			if bsearch(matrix[i], target) == true {
				return true
			}
		}
		return false
	}
}

func bsearch(data []int, target int) bool {
	start, mid, end := 0, 0, len(data)-1
	for start <= end {
		mid = (start + end) / 2
		if data[mid] == target {
			return true
		} else {
			if data[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}
