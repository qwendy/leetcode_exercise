// https://leetcode.com/problems/powx-n/description/
package practice

// func myPow(x float64, n int) float64 {
// 	result := float64(1)
// 	var a float64
// 	if x == 0 {
// 		return 0
// 	}
// 	if n > 0 {
// 		a = x
// 	} else if n < 0 {
// 		a = 1 / x
// 		n = -n
// 	} else {
// 		return 1
// 	}

// 	for n > 0 {
// 		m := 1
// 		y := a
// 		for n-m*2 >= 0 {
// 			m = m * 2
// 			y = y * y
// 		}
// 		result *= y
// 		n = n - m
// 	}
// 	return result
// }

func myPow(x float64, n int) float64 {
	result := float64(1)
	for i := n; i != 0; i = i / 2 {
		if i%2 != 0 {
			result *= x
		}
		x *= x
	}

	if n < 0 {
		result = 1 / result
	}
	return result
}
