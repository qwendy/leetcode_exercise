package leetcode

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	lo := 1
	hi := x / 2
	for {
		mid := (hi + lo) / 2
		if square(mid) <= x && square(mid+1) > x {
			return mid
		}
		if square(mid) > x {
			hi = mid - 1
		} else if square(mid) < x {
			lo = mid + 1
		}
		if lo > hi {
			return hi
		}
	}
}

func square(x int) int {
	return x * x
}
