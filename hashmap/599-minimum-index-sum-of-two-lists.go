package main

func findRestaurant(list1 []string, list2 []string) []string {
	hm := make(map[string]int)
	for idx, e := range list1 {
		hm[e] = idx
	}
	res := []string{}
	v := len(list2) * len(list1)
	for i, e := range list2 {
		if j, ok := hm[e]; ok {
			if i+j < v {
				v = i + j
				res = []string{e}
			} else if i+j == v {
				res = append(res, e)
			}
		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
