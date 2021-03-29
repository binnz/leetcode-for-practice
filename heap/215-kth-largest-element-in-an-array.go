package main

func findKthLargest(nums []int, k int) int {
	buildHeap(nums, k)
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[i], nums[0] = nums[0], nums[i]
			heapify(nums, 0, k)
		}
	}
	return nums[0]
}

func buildHeap(nums []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}
}
func heapify(nums []int, i, n int) {
	l, r := 2*i+1, 2*i+2
	minIdx := i
	if l < n && nums[l] < nums[minIdx] {
		minIdx = l
	}
	if r < n && nums[r] < nums[minIdx] {
		minIdx = r
	}
	if minIdx != i {
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
		heapify(nums, minIdx, n)
	}
}
