package main

func postorder(root *Node) []int {
	stack := []*Node{root}
	visited := []bool{false}
	res := []int{}
	if root == nil {
		return []int{}
	}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		visit := visited[len(visited)-1]
		stack = stack[:len(stack)-1]
		visited = visited[:len(visited)-1]
		if visit {
			res = append(res, cur.Val)
		} else {
			stack = append(stack, cur)
			visited = append(visited, true)
			for i := len(cur.Children) - 1; i >= 0; i-- {
				stack = append(stack, cur.Children[i])
				visited = append(visited, false)
			}
		}
	}
	return res
}
