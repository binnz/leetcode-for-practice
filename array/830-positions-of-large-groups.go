/**
https://leetcode-cn.com/problems/positions-of-large-groups/
**/
package main

func largeGroupPositions(s string) [][]int {
	var res [][]int
	var start int
	for idx := range s {
		if idx == len(s)-1 || s[idx] != s[idx+1] {
			if idx-start >= 2 {
				res = append(res, []int{start, idx})
			}
			start = idx + 1
		}
	}
	return res
}
