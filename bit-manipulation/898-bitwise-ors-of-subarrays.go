package main

func subarrayBitwiseORs(arr []int) int {
	res := make(map[int]struct{})
	cur := make(map[int]struct{})
	for _, e := range arr {
		mid := make(map[int]struct{})
		mid[e] = struct{}{}
		res[e] = struct{}{}
		for i := range cur {
			mid[e|i] = struct{}{}
			res[e|i] = struct{}{}
		}
		cur = mid
	}
	return len(res)
}
