package main

import "math/rand"

type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{origin: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	tmp := make([]int, len(this.origin))
	copy(tmp, this.origin)
	for i := len(tmp); i > 1; i-- {
		idx := rand.Intn(i)
		tmp[i-1], tmp[idx] = tmp[idx], tmp[i-1]
	}
	return tmp
}
