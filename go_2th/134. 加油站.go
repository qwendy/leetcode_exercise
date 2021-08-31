package practise

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		sum := 0
		count := 0
		for ; count < len(gas); count++ {
			index := (count + i) % len(gas)
			sum += gas[index] - cost[index]
			if sum < 0 {
				break
			}
		}
		if count == len(gas) {
			return i
		} else {
			i = i + count
		}
	}
	return -1
}
