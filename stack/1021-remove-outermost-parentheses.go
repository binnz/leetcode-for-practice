/**
https://leetcode-cn.com/problems/remove-outermost-parentheses/
**/
package main

func removeOuterParentheses(S string) string {
	var res []rune
	var count int
	for _, s := range S {

		if s == ')' {
			count--
		}

		if count >= 1 {
			res = append(res, s)
		}

		if s == '(' {
			count++
		}
	}
	return string(res)
}
