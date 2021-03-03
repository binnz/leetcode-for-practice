package main

func arrangeCoins(n int) int {
	if n <= 1 {
		return n
	}
	left, right := 0, n/2+1
	for left < right {
		mid := left + (right-left)>>1
		if mid*(mid+1)/2 == n {
			return mid
		} else if mid*(mid+1)/2 < n {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left*(left+1)/2 <= n {
		return left
	}
	return left - 1
}
