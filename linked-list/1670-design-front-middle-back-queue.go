package main

type FrontMiddleBackQueue struct {
	nums []int
}

func Constructor() FrontMiddleBackQueue {

	return FrontMiddleBackQueue{nums: make([]int, 0)}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.nums = append([]int{val}, this.nums...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	p := len(this.nums) / 2
	this.nums = append(this.nums[:p], append([]int{val}, this.nums[p:]...)...)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.nums = append(this.nums, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.nums) == 0 {
		return -1
	}
	ans := this.nums[0]
	this.nums = this.nums[1:]
	return ans
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.nums) == 0 {
		return -1
	}
	p := (len(this.nums) - 1) / 2
	ans := this.nums[p]
	this.nums = append(this.nums[:p], this.nums[p+1:]...)
	return ans

}

func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.nums) == 0 {
		return -1
	}
	ans := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return ans
}
