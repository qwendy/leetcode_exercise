// https://leetcode.com/problems/wiggle-sort-ii/description/
package practice

import (
	"fmt"
	"math/rand"
	"time"
)

func wiggleSort(nums []int) {
	median := findKthLargest(nums, len(nums)/2, 0, len(nums)-1)
	fmt.Println("median", median)
	partition3wayV(nums, 0, len(nums)-1, median)
}

// Kth smallest. quick select
func findKthLargest(nums []int, k, left, right int) int {
	if left >= right {
		return nums[right]
	}
	rand.Seed(time.Now().UnixNano())
	pivot := nums[rand.Intn(right-left+1)+left]
	smaller, larger := partition3way(nums, left, right, pivot)
	switch {
	case k-1 <= smaller:
		return findKthLargest(nums, k, left, smaller)
	case k-1 >= larger:
		return findKthLargest(nums, k, larger, right)
	default:
		return nums[k-1]
	}

}

func partition3way(nums []int, left, right, pivot int) (smaller int, larger int) {
	i, j, k := left, right, left
	for k <= j {
		if nums[k] > pivot {
			nums[k], nums[i] = nums[i], nums[k]
			i++
			k++
		} else if nums[k] < pivot {
			nums[k], nums[j] = nums[j], nums[k]
			j--
		} else {
			k++
		}
	}
	return i - 1, j + 1
}

func partition3wayV(nums []int, left, right, pivot int) (smaller int, larger int) {
	i, j, k := left, right, left
	virtualIndex := func(value int) int {
		return (value*2 + 1) % (len(nums) | 1)
	}
	for k <= j {
		fmt.Println(virtualIndex(k))
		if nums[virtualIndex(k)] > pivot {
			nums[virtualIndex(k)], nums[virtualIndex(i)] = nums[virtualIndex(i)], nums[virtualIndex(k)]
			i++
			k++
		} else if nums[virtualIndex(k)] < pivot {
			nums[virtualIndex(k)], nums[virtualIndex(j)] = nums[virtualIndex(j)], nums[virtualIndex(k)]
			j--
		} else {
			k++
		}
	}
	return i - 1, j + 1
}

func checkWiggle(nums []int) bool {
	for i := 1; i < len(nums)-1; i = i + 2 {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}
