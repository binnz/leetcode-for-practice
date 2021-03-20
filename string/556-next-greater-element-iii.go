package main

import "fmt"
import "sort"

func nextGreaterElement(n int) int {
	nums := []int{}
	for n > 0 {
		c := n % 10
		n /= 10
		nums = append([]int{c}, nums...)
	}
	fmt.Println(nums)
	if len(nums) == 1 {
		return -1
	}
	max := true
	var ttt int
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			max = false
			ttt = i
			break
		}
	}
	if max {
		return -1
	}
	tmp := nums[ttt+1:]
	sort.Ints(tmp)
	target := 0
	for target < len(tmp) {
		if tmp[target] > nums[ttt] {
			break
		}
		target++
	}
	t := nums[ttt]
	nums = nums[:ttt]
	nums = append(nums, tmp[target])
	tmp[target] = t
	sort.Ints(tmp)
	nums = append(nums, tmp...)
	res := 0
	for i, v := range nums {
		res += v
		if i != len(nums)-1 {
			res *= 10
		}
	}
	if res >= 1<<31 {
		return -1
	}
	return res
}
