package main

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	left, right := 1, x/2+1
	for left < right {
		mid := left + (right-left)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}
