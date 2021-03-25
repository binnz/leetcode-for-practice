package main

func lexicalOrder(n int) []int {
	res := []int{}

	for i := 1; i < 10; i++ {
		if i > n {
			break
		}
		stack := []int{i}
		for len(stack) != 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur)
			for i := 9; i >= 0; i-- {
				if cur*10+i <= n {
					stack = append(stack, cur*10+i)
				}
			}
		}
	}
	return res
}
