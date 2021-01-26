package main

func maxTurbulenceSize(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	res, start := 1, 0
	i, prev := 1, arr[0]-arr[1]
	for i < len(arr) {
		if (arr[i]-arr[i-1])*prev < 0 {
			if i-start+1 > res {
				res = i - start + 1
			}
		} else {
			start = i - 1
		}
		prev = arr[i] - arr[i-1]
		i++
	}
	return res
}
