// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/32/
package easy_string

func reverseString(s string) string {
	n := len(s)
	r := []byte(s)
	for i := 0; i < n/2; i++ {
		tmp := r[i]
		r[i] = r[n-i-1]
		r[n-i-1] = tmp
	}
	return string(r)
}
