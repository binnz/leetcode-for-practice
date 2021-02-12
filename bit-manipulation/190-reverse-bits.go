package main

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1
		if num&1 == 1 {
			res |= 1
		}

		num >>= 1
	}
	return res
}
