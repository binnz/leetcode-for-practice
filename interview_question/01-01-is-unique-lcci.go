package main

func isUnique(astr string) bool {
	if len(astr) > 26 {
		return false
	}
	bits := make([]int, 26)
	for _, e := range astr {
		if bits[e-'a'] == 1 {
			return false
		} else {
			bits[e-'a'] = 1
		}
	}
	return true
}
