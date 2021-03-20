/**
https://leetcode-cn.com/problems/fibonacci-number/
**/
package main

func fib(n int) int {
	a, b, c := 0, 1, 1
	if n == 0 {
		return 0
	}
	for i := 2; i < n; i++ {
		a = b
		b = c
		c = b + a

	}
	return c
}
