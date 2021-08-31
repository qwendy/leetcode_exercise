package practise

func canCompleteCircuit(gas []int, cost []int) int {
	start := -1
	for i := 0; i < 2*len(gas)-1; i++ {
		sum := 0
		j := i
		for ; j < len(gas)+i; j++ {
			index := j % len(gas)
			sum += gas[index] - cost[index]
			if sum < 0 {
				break
			}
		}
	}
	return start
}
