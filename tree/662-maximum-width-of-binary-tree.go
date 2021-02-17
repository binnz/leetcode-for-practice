package main

func widthOfBinaryTree(root *TreeNode) int {
	left := []int{0}
	res := 0
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level int, index int) {
		if node == nil {
			return
		}
		if level > len(left)-1 {
			left = append(left, index)
		}
		if index-left[level]+1 > res {
			res = index - left[level] + 1
		}
		dfs(node.Left, level+1, 2*index-1)
		dfs(node.Right, level+1, 2*index)
	}
	dfs(root, 0, 0)
	return res
}

func widthOfBinaryTree2(root *TreeNode) int {
	var (
		stack []*TreeNode = []*TreeNode{root}
		mid   []*TreeNode
		res   int
		minV  int
		maxV  int
		hm    map[*TreeNode]int = map[*TreeNode]int{}
	)
	hm[root] = 0
	for len(stack) != 0 {
		for i := 0; i < len(stack); i++ {
			v := hm[stack[i]]
			e := stack[i]
			if i == 0 {
				minV = v
			}
			maxV = v
			if e.Left != nil {
				mid = append(mid, e.Left)
				hm[e.Left] = 2*v - 1
			}
			if e.Right != nil {
				mid = append(mid, e.Right)
				hm[e.Right] = 2 * v
			}
		}
		res = max(res, maxV-minV+1)
		stack = mid
		mid = []*TreeNode{}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
