package main

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{longer * k}
	}
	res := make([]int, k+1)
	for i := 0; i <= k; i++ {
		res[i] = k*shorter + (longer-shorter)*i
	}
	return res
}
