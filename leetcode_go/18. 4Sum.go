// https://leetcode.com/problems/4sum/description/
package practice

import (
	"fmt"
	"sort"
	"strconv"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return make([][]int, 0)
	}
	// sort
	sort.Ints(nums)
	fmt.Println(nums)
	// iterate
	n := len(nums)
	r := make([][]int, 0)
	m := make(map[string]bool, 0)
	for h := 0; h < n-3; h++ {
		if h > 0 && nums[h] == nums[h-1] {
			continue
		}
		for i := h + 1; i < n-2; i++ {
			if i > h+1 && nums[i] == nums[i-1] {
				continue
			}
			sum := target - nums[h] - nums[i]
			j := i + 1
			k := n - 1
			for j < k {
				switch {
				case sum < nums[j]+nums[k]:
					k--
				case sum > nums[j]+nums[k]:
					j++
				default:
					tmp := []int{nums[h], nums[i], nums[j], nums[k]}
					s := strconv.Itoa(tmp[0]) + "," + strconv.Itoa(tmp[1]) + "," + strconv.Itoa(tmp[2]) + "," + strconv.Itoa(tmp[3])
					if _, ok := m[s]; !ok {
						m[s] = true
						r = append(r, tmp)
					}
					j++
					k--
				}
			}
		}
	}
	return r
}
