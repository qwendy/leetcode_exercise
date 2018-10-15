// https://leetcode.com/problems/wiggle-sort-ii/description/
package leetcode

import (
	"fmt"
	"math/rand"
	"time"
)

func wiggleSort(nums []int) {

}

// Kth smallest
func findKthSmallest(nums []int, k, left, right int) int {
	if left >= right {
		return nums[right]
	}
	rand.Seed(time.Now().UnixNano())
	median := nums[rand.Intn(right-left+1)+left]
	smaller, larger := threeWayPartition(nums, left, right, median)
	fmt.Println(smaller, larger, median, left, right, nums)
	switch {
	case k-1 <= smaller:
		findKthSmallest(nums, k, left, smaller)
	case k-1 >= larger:
		findKthSmallest(nums, k, larger, right)
	default:
		return nums[k-1]
	}
	return nums[0]
}

func threeWayPartition(nums []int, left, right, median int) (smaller int, larger int) {
	i, j, k := left, right, left
	for k <= right && k <= j {
		if nums[k] < median {
			nums[k], nums[i] = nums[i], nums[k]
			i++
			k++
		} else if nums[k] > median {
			nums[k], nums[j] = nums[j], nums[k]
			j--
		} else {
			k++
		}
	}
	return i - 1, j + 1
}
