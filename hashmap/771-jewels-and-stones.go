package main

func numJewelsInStones(J string, S string) int {
	hm := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		hm[S[i]]++
	}
	res := 0
	for i := 0; i < len(J); i++ {
		if v, ok := hm[J[i]]; ok {
			res += v
		}
	}
	return res
}
