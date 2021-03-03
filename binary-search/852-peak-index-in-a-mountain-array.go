package main

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == len(arr) {
		return left - 1
	}
	return left
}
