package main

func findMaximumXOR(nums []int) int {
	res := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		mask |= (1 << i)
		set := map[int]struct{}{}
		tmp := res | (1 << i)
		for _, n := range nums {
			set[mask&n] = struct{}{}
			if _, ok := set[tmp^(mask&n)]; ok {
				res = tmp
				break
			}
		}
	}
	return res
}
