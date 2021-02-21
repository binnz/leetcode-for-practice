package main

func findKthPositive(arr []int, k int) int {

	if arr[0] > k {
		return k
	}
	if arr[len(arr)-1] == len(arr) {
		return len(arr) + k
	}
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)>>1
		if arr[mid] >= mid+k+1 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return k + l
}
