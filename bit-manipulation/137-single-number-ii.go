package main

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < 64; i++ {
		c := 0
		for _, e := range nums {
			c += e >> i & 1
		}
		res |= c % 3 << i
	}
	return res
}
