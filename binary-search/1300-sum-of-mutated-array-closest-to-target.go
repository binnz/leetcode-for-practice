package main

func findBestValue(arr []int, target int) int {

	left, right := 0, 100000

	for left < right {
		mid := left + (right-left)>>1
		if calSum(arr, mid) < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if right == 10000 {
		res := 0
		for _, e := range arr {
			if e > res {
				res = e
			}
		}
		return res
	}
	if target-calSum(arr, left-1) <= calSum(arr, left)-target {
		return left - 1
	}
	return left

}

func calSum(arr []int, mid int) int {
	res := 0
	for _, e := range arr {
		res += min(e, mid)
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	findBestValue([]int{2, 3, 5}, 11)
}
