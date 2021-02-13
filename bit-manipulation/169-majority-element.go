package main

func majorityElement(nums []int) int {
	var res int
	c := 0
	for e := range nums {
		if c == 0 {
			res = e
			c++
		} else {
			if e == res {
				c++
			} else {
				c--
			}
		}
	}
	return res
}

func main() {
	majorityElement([]int{3, 2, 3})
}
