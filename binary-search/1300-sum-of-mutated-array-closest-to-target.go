package main

import "sort"

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	left, right := 0, len(arr)
	t := float64(target) / float64(len(arr))
	for left < right {
		mid := left + (right-left)>>1
		if float64(arr[mid]) == t {
			left = mid
			break
		} else if float64(arr[mid]) < t {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == 0 {
		return int(t)
	}
	if left == len(arr) {
		return arr[len(arr)-1]
	}
	if abs(float64(target), 5*float64(arr[left])) > abs(float64(target), 5*float64(arr[left-1])) {
		return arr[left-1]
	} else {
		return arr[left]
	}

}
func abs(a, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	findBestValue([]int{2, 3, 5}, 10)
}
