package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	mid, res := make(map[int]struct{}), make(map[string]struct{})
	str := 0
	for i := 0; i < 10; i++ {
		str <<= 3
		str |= int(s[i]) & 7

	}
	mid[str] = struct{}{}
	for i := 10; i < len(s); i++ {
		str <<= 3
		str |= int(s[i]) & 7
		str &= 0x3fffffff
		if _, ok := mid[str]; ok {
			res[s[i-9:i+1]] = struct{}{}
		} else {
			mid[str] = struct{}{}
		}
	}
	output := []string{}
	for k := range res {
		output = append(output, k)
	}
	return output
}
