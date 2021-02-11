package main

func hasAlternatingBits(n int) bool {
	i := n ^ (n >> 1) + 1
	if i&(i-1) == 0 {
		return true
	} else {
		return false
	}
}
