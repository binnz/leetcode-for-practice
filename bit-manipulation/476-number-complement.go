package main

func findComplement(num int) int {
	res, i := 0, 0
	for num > 0 {
		if num&1 == 0 {
			res |= (1 << i)
		}
		num >>= 1
		i++
	}
	return res
}
