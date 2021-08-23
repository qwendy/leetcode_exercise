// https://leetcode.com/problems/3sum/description/

package practice

import (
	"fmt"
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
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
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		sum := -nums[i]
		j := i + 1
		k := n - 1
		for j < k {

			switch {
			case sum < nums[j]+nums[k]:
				k--
			case sum > nums[j]+nums[k]:
				j++
			default:
				tmp := []int{nums[i], nums[j], nums[k]}
				s := strconv.Itoa(tmp[0]) + "," + strconv.Itoa(tmp[1]) + "," + strconv.Itoa(tmp[2])
				if _, ok := m[s]; !ok {
					m[s] = true
					r = append(r, tmp)
				}
				j++
				k--
			}
		}
	}
	return r
}
