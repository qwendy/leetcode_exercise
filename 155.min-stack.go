/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (39.41%)
 * Likes:    2248
 * Dislikes: 243
 * Total Accepted:    368.1K
 * Total Submissions: 931.8K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 *
 *
 *
 *
 * Example:
 *
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 *
 *
 *
 *
 */
package practice

// @lc code=start
type MinStack struct {
	Count     int
	Min       int
	Container []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Container: make([]int, 100),
	}
}

func (this *MinStack) Push(x int) {
	if this.Count >= len(this.Container) {
		temp := make([]int, this.Count*2)
		copy(temp, this.Container)
		this.Container = temp
	}
	this.Count++
	this.Container[this.Count-1] = x
	if this.Count == 1 || x < this.Min {
		this.Min = x
	}
}

func (this *MinStack) Pop() {
	if this.Count <= 0 {
		return
	}
	this.Count--
	if this.Container[this.Count] == this.Min {
		this.Min = this.Container[0]
		for i := 1; i < this.Count; i++ {
			if this.Container[i] < this.Min {
				this.Min = this.Container[i]
			}
		}
	}
}

func (this *MinStack) Top() int {
	if this.Count > 0 {
		return this.Container[this.Count-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	return this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
