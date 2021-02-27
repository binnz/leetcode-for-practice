package main

func canPermutePalindrome(s string) bool {
	hm := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		hm[s[i]]++
	}
	c := 0
	for _, v := range hm {
		if v%2 == 1 {
			c++
		}
		if c > 1 {
			return false
		}
	}
	return true
}
