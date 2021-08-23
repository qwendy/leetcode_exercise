// https://leetcode.com/problems/search-a-2d-matrix/description/
package practice

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	r := 0
	for r+1 < m && len(matrix[r+1]) > 0 && matrix[r+1][0] <= target {
		r++
	}
	lo, hi := 0, len(matrix[r])-1
	for lo <= hi {
		mid := (hi + lo) / 2
		middle := matrix[r][mid]
		if middle < target {
			lo = mid + 1
		} else if middle > target {
			hi = mid - 1
		} else {
			return true
		}
	}
	return false
}
