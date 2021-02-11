package main

func countPrimeSetBits(L int, R int) int {
	countInt := func(i int) int {
		c := 0
		for i > 0 {
			if i&1 == 1 {
				c++
			}
			i >>= 1
		}
		return c
	}
	res := 0
	for i := L; i <= R; i++ {
		switch countInt(i) {
		case 2, 3, 5, 7, 11, 13, 17, 19:
			res++
		default:
			continue
		}
	}
	return res
}
