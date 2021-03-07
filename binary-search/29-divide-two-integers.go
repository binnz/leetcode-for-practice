package main

func divide(dividend int, divisor int) int {
	var diff bool
	if dividend == -1<<31 && divisor == -1 {
		return (1 << 31) - 1
	}
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		diff = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	res := 0
	for divisor <= dividend {
		i := 0
		for divisor<<(i+1) < dividend {
			i++
		}
		res += 1 << i
		dividend -= divisor << i
	}
	if diff {
		return -res
	}
	return res
}
