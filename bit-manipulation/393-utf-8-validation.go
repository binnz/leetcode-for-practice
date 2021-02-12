package main

func validUtf8(data []int) bool {
	length, i := len(data), 0

	for i < length {
		if data[i]>>7&1 == 0 {
			i++
			continue
		}
		clen, s := 0, 7
		for s > 3 {
			if data[i]>>s&1 == 1 {
				clen++
				s--
			} else {
				break
			}
		}
		if data[i]>>s&1 == 1 {
			return false
		}
		if clen < 2 {
			return false
		}
		if i+clen > length {
			return false
		}
		i++
		j := i
		for i < j+clen-1 {
			if data[i]>>7&1 == 0 {
				return false
			}
			if data[i]>>6&1 == 1 {
				return false
			}
			i++
		}

	}
	return true
}

func validUtf82(data []int) bool {
	count := func(n int) int {
		if n>>3 == 0b11110 {
			return 4
		} else if n>>4 == 0b1110 {
			return 3
		} else if n>>5 == 0b110 {
			return 2
		} else if n>>7 == 0b0 {
			return 1
		} else {
			return -1
		}
	}
	clen := 0
	for _, e := range data {
		if clen == 0 {
			clen = count(e)
			if clen < 0 {
				return false
			}
		} else {
			if !(e>>6 == 0b10) {
				return false
			}
		}
		clen--
	}
	return clen == 0
}
