package main

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	i, s := 0, 0
	for s < len(s1) {
		if s1[i] != s2[s] {
			s++
		} else {
			j := s
			for i < len(s1) && s1[i] == s2[j] {
				i++
				j++
				if i < len(s1) && j == len(s2) {
					j = 0
				}
			}
			if i == len(s1) && s1[i-1] == s2[j-1] {
				return true
			}
			i = 0
			s++
			continue
		}
	}
	return false
}
