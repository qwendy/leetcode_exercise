// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/31/
package array

func m_rotate(matrix [][]int) {
	n := len(matrix)
	// (i, j) -> (j, n-i-1)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			var replaceCoordinate = func(value, i, j int) (int, int, int) {
				r, c := j, n-i-1
				tmp := matrix[r][c]
				matrix[r][c] = value
				return tmp, r, c
			}
			tmp := matrix[i][j]
			r := i
			c := j
			for k := 0; k < 4; k++ {
				tmp, r, c = replaceCoordinate(tmp, r, c)
			}
		}
	}
}
