package main

func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); {
		if i == len(bits)-1 {
			return true
		} else if bits[i] == 1 {
			i = i + 2
		} else {
			i++
		}
	}
	return false
}
