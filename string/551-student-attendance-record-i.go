package main

func checkRecord(s string) bool {
	ac := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "A" {
			ac++
			if ac > 1 {
				return false
			}
		}

		if i < len(s)-2 && string(s[i]) == "L" && string(s[i+1]) == "L" && string(s[i+2]) == "L" {
			return false
		}
	}
	return true
}
