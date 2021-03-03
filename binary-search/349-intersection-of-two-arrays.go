package main

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	hm := make(map[int]struct{}, len(nums1))
	for _, e := range nums1 {
		hm[e] = struct{}{}
	}
	res := map[int]struct{}{}
	for _, e := range nums2 {
		if _, ok := hm[e]; ok {
			res[e] = struct{}{}
		}
	}
	ress := []int{}
	for k, _ := range res {
		ress = append(ress, k)
	}
	return ress
}
