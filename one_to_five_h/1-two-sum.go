/*
    Leetcode 1
	https://leetcode-cn.com/problems/two-sum/

*/
package twoSum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums{
		if idx, isOk := m[target-val]; isOk{
			return []int{idx, index}
		}
		m[val] = index
	}
	return nil
}
