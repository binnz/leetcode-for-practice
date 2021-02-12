package main

func singleNumber(nums []int) []int {
	mid := 0
	for _, e := range nums {
		mid ^= e
	}
	diff := mid & (^mid + 1)
	res := []int{0, 0}
	for _, e := range nums {
		if e&diff == 0 {
			res[0] ^= e
		} else {
			res[1] ^= e
		}
	}
	return res
}
