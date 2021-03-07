package main

func myPow(x float64, n int) float64 {
	var dfs func(float64, int) float64
	hm := map[int]float64{}
	dfs = func(x float64, n int) float64 {
		if e, ok := hm[n]; ok {
			return e
		}
		if n == 1 {
			return x
		}
		if n == 0 {
			return 1
		}
		negative := false
		if n < 0 {
			negative = true
			n = -n
		}
		c := n / 2
		p := n % 2
		var res float64
		if p == 1 {
			res = dfs(x, c) * dfs(x, c) * x
		} else {
			res = dfs(x, c) * dfs(x, c)
		}
		if negative {
			hm[n] = 1 / res
			return 1 / res
		}
		hm[n] = res
		return res
	}
	return dfs(x, n)
}

func myPow1(x float64, n int) float64 {
	var res float64 = 1.0
	var ng bool
	if n < 0 {
		ng = true
		n = -n
	}
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	if ng {
		return 1 / res
	}
	return res
}
