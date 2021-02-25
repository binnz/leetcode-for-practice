package main

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false

	}
	hm := make(map[byte]int, len(s1))
	for _, e := range s1 {
		hm[byte(e)]++

	}
	for _, e := range s2 {
		if v, ok := hm[byte(e)]; ok && v > 0 {
			hm[byte(e)]--

		} else {
			return false

		}

	}
	return true

}
