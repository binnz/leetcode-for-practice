package main

func maximumSwap(num int) int {
	stack := []int{}
	if num/10 == 0 {
		return num
	}
	for num > 0 {
		v := num % 10
		num /= 10
		stack = append([]int{v}, stack...)
	}
	for i := 0; i < len(stack); i++ {
		maxIdx := i
		maxV := stack[i]
		for j := i; j < len(stack); j++ {
			if stack[j] >= maxV {
				maxV = stack[j]
				maxIdx = j
			}
		}
		if maxIdx != i && stack[maxIdx] != stack[i] {
			stack[maxIdx], stack[i] = stack[i], stack[maxIdx]
			break
		}
	}
	res := 0
	for idx, e := range stack {
		res += e
		if idx != len(stack)-1 {
			res *= 10
		}

	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
