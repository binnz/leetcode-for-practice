package main

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1
	for right-left >= k {
		if arr[right]-x >= x-arr[left] {
			right--
		} else {
			left++
		}
	}
	return arr[left : right+1]
}
func findClosestElements1(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	for left < right {
		mid := left + (right-left)>>1
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}
