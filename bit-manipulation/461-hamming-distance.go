package main

func hammingDistance(x int, y int) int {
	res := 0
	for i := 0; i < 31; i++ {
		if (y & (1 << i)) != (x & (1 << i)) {
			res++
		}
	}
	return res
}
