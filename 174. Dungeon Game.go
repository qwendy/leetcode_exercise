package practice

import (
	"math"
)

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	hp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		hp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			hp[i][j] = math.MaxInt32
		}
	}
	hp[m][n-1] = 1
	hp[m-1][n] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			min := hp[i+1][j]
			d := hp[i][j+1]
			if min > d {
				min = d
			}
			need := min - dungeon[i][j]
			if need <= 0 {
				hp[i][j] = 1
			} else {
				hp[i][j] = need
			}
		}
	}
	return hp[0][0]

}
