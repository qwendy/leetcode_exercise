package practise

import "fmt"

//拼接最大数
////给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n) 个数字拼接
//成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
//
// 求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
//
// 说明: 请尽可能地优化你算法的时间和空间复杂度。
//
// 示例 1:
//
// 输入:
//nums1 = [3, 4, 6, 5]
//nums2 = [9, 1, 2, 5, 8, 3]
//k = 5
//输出:
//[9, 8, 6, 5, 3]
//
// 示例 2:
//
// 输入:
//nums1 = [6, 7]
//nums2 = [6, 0, 4]
//k = 5
//输出:
//[6, 7, 6, 0, 4]
//
// 示例 3:
//
// 输入:
//nums1 = [3, 9]
//nums2 = [8, 9]
//k = 3
//输出:
//[9, 8, 9]
// Related Topics 栈 贪心 单调栈 👍 409 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	max := []int{}
	for i := 0; i <= k; i++ {
		if i > len(nums1) {
			continue
		}
		if k-i > len(nums2) {
			continue
		}

		sub1 := maxSub(nums1, i)
		sub2 := maxSub(nums2, k-i)
		m := merge(sub1, sub2)
		fmt.Println(i, sub1, sub2, m)
		if len(max) == 0 {
			max = m
		} else {
			if bigger(m, max) {
				max = m
			}
		}
	}
	return max
}
func maxSub(a []int, k int) []int {
	stack := []int{}
	drop := len(a) - k
	for i, _ := range a {
		for len(stack) > 0 && stack[len(stack)-1] < a[i] {
			if drop > 0 {
				stack = stack[:len(stack)-1]
				drop--
			} else {
				break
			}
		}
		stack = append(stack, a[i])
		if drop == 0 && len(stack) >= k {
			return stack[:k]
		}
	}
	fmt.Println(a, k, stack)
	return stack[:k]
}

func bigger(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true
}
func merge(a, b []int) []int {
	result := make([]int, 0)
	indexA, indexB := 0, 0
	for indexA < len(a) && indexB < len(b) {
		if a[indexA] > b[indexB] {
			result = append(result, a[indexA])
			indexA++
		} else if a[indexA] < indexB {
			result = append(result, b[indexB])
			indexB++
		} else {
			result = append(result, a[indexB])
			result = append(result, b[indexB])
			indexA++
			indexB++
		}

	}
	result = append(result, a[indexA:]...)
	result = append(result, b[indexB:]...)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

// 2021-08-30 18:37:01
