package main

import "math"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	res := []byte{}
	var n uint32 = uint32(num)
	for n > 0 {
		var i byte
		if i = byte(n & 0xf); i < 10 {
			i += 48
		} else {
			i += 87
		}
		res = append([]byte{i}, res...)
		n >>= 4
	}
	return string(res)
}

func toHex2(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num += int(math.Pow(2, 32))
	}
	res := []byte{}
	hash := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

	for num > 0 {
		tmp := num % 16
		num /= 16
		res = append([]byte{hash[tmp]}, res...)
	}
	return string(res)
}
