package practise

//Candy
////There are n children standing in a line. Each child is assigned a rating value
// given in the integer array ratings.
//
// You are giving candies to these children subjected to the following requireme
//nts:
//
//
// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
//
//
// Return the minimum number of candies you need to have to distribute the candi
//es to the children.
//
//
// Example 1:
//
//
//Input: ratings = [1,0,2]
//Output: 5
//Explanation: You can allocate to the first, second and third child with 2, 1,
//2 candies respectively.
//
//
// Example 2:
//
//
//Input: ratings = [1,2,2]
//Output: 4
//Explanation: You can allocate to the first, second and third child with 1, 2,
//1 candies respectively.
//The third child gets 1 candy because it satisfies the above two conditions.
//
//
//
// Constraints:
//
//
// n == ratings.length
// 1 <= n <= 2 * 104
// 0 <= ratings[i] <= 2 * 104
//
// Related Topics Array Greedy
// 👍 2129 👎 223

//leetcode submit region begin(Prohibit modification and deletion)
func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}

	candy := make([]int, len(ratings))
	for i, _ := range candy {
		candy[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candy[i-1] <= candy[i] {
			candy[i-1] = candy[i] + 1
		}
	}

	count := 0
	for _, v := range candy {
		count += v
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
