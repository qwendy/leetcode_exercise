package practice

// Determine whether an integer is a palindrome. Do this without extra space.

func isPalindrome(x int) bool {
	y, t := 0, x
	for t > 0 {
		y = y*10 + t%10
		t = t / 10
	}
	return x == y
}
