package main

func findUnsortedSubarray(arr []int) int {
	l, r := 0, len(arr)-1
	if len(arr) < 2 {
		return 0
	}
	for l < r {
		if arr[l] <= arr[l+1] {
			l++
		} else {
			break
		}
	}
	for l < r {
		if arr[r-1] <= arr[r] {
			r--
		} else {
			break
		}
	}
	if l == r {
		return 0
	}
	l--
	r++
	for l >= 0 || r < len(arr) {
		minV, maxV := maxMin1(arr, l+1, r-1)
		b := true
		if l >= 0 && minV < arr[l] {
			l--
			b = false
		}
		if r < len(arr) && maxV > arr[r] {
			r++
			b = false
		}
		if b {
			break
		}
	}
	return r - l - 1
}
func maxMin1(arr []int, l, r int) (int, int) {
	minV := 2 << 31
	maxV := -2 << 31
	for i := l; i <= r; i++ {
		if arr[i] > maxV {
			maxV = arr[i]
		}
		if arr[i] < minV {
			minV = arr[i]
		}
	}
	return minV, maxV
}
