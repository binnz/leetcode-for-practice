package main

func pancakeSort(arr []int) []int {
	res := []int{}
	for j := len(arr) - 1; j >= 0; j-- {
		idx := maxIndex(arr, j)
		res = append(res, idx+1)
		reverse(arr, 0, idx)
		res = append(res, j+1)
		reverse(arr, 0, j)
	}
	return res
}

func reverse(arr []int, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func maxIndex(arr []int, idx int) int {
	maxIdx := 0
	for i := 0; i <= idx; i++ {
		if arr[i] > arr[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}
