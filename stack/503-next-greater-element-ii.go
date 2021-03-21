package main

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = -1
	}
	stack := []int{}
	n := len(nums)
	for i := 0; i < 2*len(nums)-1; i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return res
}

func main() {
	nextGreaterElements([]int{1, 2, 3, 4, 3})
}
