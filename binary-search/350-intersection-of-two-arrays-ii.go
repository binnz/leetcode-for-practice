package main

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	hm := make(map[int]int, len(nums1))
	for _, e := range nums1 {
		hm[e]++
	}
	res := []int{}
	for _, e := range nums2 {
		if c, ok := hm[e]; ok && c > 0 {
			res = append(res, e)
			hm[e]--
		}
	}
	return res
}
