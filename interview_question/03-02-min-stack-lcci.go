package main

type MinStack struct {
	stack  []int
	minVal []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minVal) == 0 || (len(this.minVal) != 0 && x <= this.GetMin()) {
		this.minVal = append(this.minVal, x)
	}
}

func (this *MinStack) Pop() {
	var min int
	if len(this.stack) > 0 {
		min = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	if min == this.GetMin() {
		this.minVal = this.minVal[:len(this.minVal)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.minVal) > 0 {
		return this.minVal[len(this.minVal)-1]
	}
	return -1
}
