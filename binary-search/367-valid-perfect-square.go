package main

func isPerfectSquare(num int) bool {
	left, right := 1, num/2
	for left < right {
		mid := left + (right-left)>>1
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left*left == num
}
