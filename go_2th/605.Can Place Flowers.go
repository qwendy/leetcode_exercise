package practise

//Can Place Flowers
//can-place-flowers
//605
////You have a long flowerbed in which some of the plots are planted, and some are
// not. However, flowers cannot be planted in adjacent plots.
//
// Given an integer array flowerbed containing 0's and 1's, where 0 means empty
//and 1 means not empty, and an integer n, return if n new flowers can be planted
//in the flowerbed without violating the no-adjacent-flowers rule.
//
//
// Example 1:
// Input: flowerbed = [1,0,0,0,1], n = 1
//Output: true
// Example 2:
// Input: flowerbed = [1,0,0,0,1], n = 2
//Output: false
//
//
// Constraints:
//
//
// 1 <= flowerbed.length <= 2 * 104
// flowerbed[i] is 0 or 1.
// There are no two adjacent flowers in flowerbed.
// 0 <= n <= flowerbed.length
//
// Related Topics Array Greedy
// 👍 1706 👎 512

//leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := len(flowerbed)
	if count < n || count < 1 {
		return false
	}
	if count == 1 && flowerbed[0] == 0 && n == 1 {
		return true
	}
	for i := 0; i < count && n >= 0; i++ {
		if flowerbed[i] == 0 &&
			(i-1 < 0 || flowerbed[i-1] == 0) &&
			(i+1 >= count || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}

//leetcode submit region end(Prohibit modification and deletion)

//2021-08-25 18:16:31
