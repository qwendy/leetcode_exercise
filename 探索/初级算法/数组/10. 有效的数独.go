// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/30/

package array

func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]bool, 9)
	column := make([]map[byte]bool, 9)
	matrix := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			if i == 0 {
				column[j] = make(map[byte]bool, 9)
			}
			matrixIndex := (j/3)*3 + i/3
			if i%3 == 0 && j%3 == 0 {
				matrix[matrixIndex] = make(map[byte]bool, 9)
			}
			v := board[i][j]
			if v == '.' {
				continue
			}
			if row[i][v] || column[j][v] || matrix[matrixIndex][v] {
				return false
			}
			row[i][v], column[j][v], matrix[matrixIndex][v] = true, true, true
		}
	}
	return true
}
