// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/22/

package array

func maxProfit(prices []int) int {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		i = findsmallest(i, prices)
		if i == -1 {
			return sum
		}
		biggerIndex := findbigger(i, prices)
		if biggerIndex != -1 {
			sum += prices[biggerIndex] - prices[i]
			i = biggerIndex
		}
	}
	return sum
}

func findbigger(index int, prices []int) int {
	for i := index + 1; i < len(prices); i++ {
		if prices[i] > prices[index] {
			for ; i < len(prices)-1; i++ {
				if prices[i] > prices[i+1] {
					return i
				}
			}
			return len(prices) - 1
		}
	}
	return -1
}

func findsmallest(index int, prices []int) int {
	for i := index; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			return i
		}
	}
	return -1
}
