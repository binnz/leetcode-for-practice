package main

func validPalindrome(s string) bool {
	if s == "" {
		return true
	}
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			if isPalindroem(s[l:r]) || isPalindroem(s[l+1:r+1]) {
				return true
			}
			return false
		}
	}
	return true
}

func isPalindroem(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
