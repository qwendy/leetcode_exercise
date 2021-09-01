package practise

import "fmt"

func compareVersion(version1 string, version2 string) int {
	index1, index2 := 0, 0
	for index1 < len(version1) && index2 < len(version2) {
		for ; index1 < len(version1); index1++ {
			if version1[index1] != '0' {
				break
			}
		}
		for ; index2 < len(version2); index2++ {
			if version2[index2] != '0' {
				break
			}
		}
		result := 0
		for {
			finished1 := false
			if index1 >= len(version1) || version1[index1] == '.' {
				finished1 = true
			}
			finished2 := false
			if index2 >= len(version2) || version2[index2] == '.' {
				finished2 = true
			}
			if finished1 && finished2 {
				if result != 0 {
					return result
				}
				if index2 < len(version2) {
					index2++
				}
				if index1 < len(version1) {
					index1++
				}
				break
			}
			if finished1 {
				return -1
			}
			if finished2 {
				return 1
			}
			if version1[index1] > version2[index2] {
				result = 1
			} else if version1[index1] < version2[index2] {
				result = -1
			}
			index1++
			index2++

		}
	}
	if index1 >= len(version1) && index2 >= len(version2) {
		return 0
	}
	if index1 < len(version1) {
		if zero(version1[index1:]) {
			return 0
		}
		return 1
	}
	if zero(version2[index2:]) {
		return 0
	}
	return -1
}

func zero(s string) bool {
	fmt.Println(s)
	for i, _ := range s {
		if s[i] != '0' && s[i] != '.' {
			return false
		}
	}
	return true
}
