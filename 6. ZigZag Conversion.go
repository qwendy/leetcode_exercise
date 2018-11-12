package practice

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string text, int nRows);
// convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows <= 1 {
		return s
	}
	rows := make([][]byte, numRows)
	r := 0
	direction := 1
	result := []byte(s)
	for _, b := range result {
		rows[r] = append(rows[r], b)
		if r == 0 {
			direction = 1
		}
		if r == numRows-1 {
			direction = -1
		}
		r = r + direction
	}
	i := 0
	for _, row := range rows {
		for _, r := range row {
			result[i] = r
			i++
		}
	}
	return string(result)
}
