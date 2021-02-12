package main

func integerReplacement(n int) int {
	if n < 3 {
		return n - 1
	}
	if n&1 == 0 {
		return integerReplacement(n>>1) + 1
	} else {
		return min(integerReplacement(n-1), integerReplacement(n+1)) + 1
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func integerReplacement(n int) int {
	res := 0
	for n > 1 {
		if n <= 3 {
			return res + n - 1
		}
		if n%2 == 0 {
			n /= 2
		} else {
			a := maxPower(n + 1)
			b := maxPower(n - 1)
			if a > b {
				n++
			} else {
				n--
			}
		}
		res++
	}
	return res
}

func maxPower(n int) int {
	i := 0
	for n%2 == 0 {
		i++
		n /= 2
	}
	return i
}
