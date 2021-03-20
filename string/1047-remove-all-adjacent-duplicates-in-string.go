/**
https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/
**/
package main

func removeDuplicates(S string) string {
	var res []rune
	for _, s := range S {
		if len(res) != 0 && res[len(res)-1] == s {
			res = res[:len(res)-1]
		} else {
			res = append(res, s)
		}
	}
	return string(res)
}
