package main

func shipWithinDays(weights []int, D int) int {
	var (
		left  int
		right int
	)
	for _, e := range weights {
		right += e
		left = max(left, e)
	}
	for left < right {
		sum := 0
		day := 0
		mid := left + (right-left)>>1
		for _, e := range weights {
			sum += e
			if sum > mid {
				day++
				sum = e
			}
		}
		if sum <= mid {
			day++
		}
		if day > D {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
