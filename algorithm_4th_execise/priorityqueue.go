package practice

type priorityqueue struct {
	nums  []int
	count int
}

func NewPriorityQueue() *priorityqueue {
	return &priorityqueue{
		nums:  make([]int, 10),
		count: 0,
	}
}

func (pq *priorityqueue) insert(a int) {
	if pq.count == len(pq.nums) {
		pq.nums = append(pq.nums, make([]int, 10)...)
	}
	pq.count++
	pq.nums[pq.count] = a
	pq.swim(pq.count)
}

func (pq *priorityqueue) delMax() int {
	max := pq.nums[1]
	pq.nums[1] = pq.nums[pq.count]
	pq.count--
	pq.sink(1)
	return max
}

func (pq *priorityqueue) swim(k int) {
	for i := k; pq.nums[i] > pq.nums[i/2] && i > 1; i /= 2 {
		pq.nums[i], pq.nums[i/2] = pq.nums[i/2], pq.nums[i]
	}
}
func (pq *priorityqueue) sink(k int) {
	for 2*k <= pq.count {
		j := 2 * k
		if j < pq.count && pq.nums[j] < pq.nums[j+1] {
			j++
		}
		if pq.nums[k] > pq.nums[j] {
			break
		}
		pq.nums[k], pq.nums[j] = pq.nums[j], pq.nums[k]
		k = j
	}

}
