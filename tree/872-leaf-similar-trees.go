/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root2 == nil || root2 == nil {
		return false
	}
	r1 := inorder(root1)
	r2 := inorder(root2)
	i, ib := r1()
	j, jb := r2()
	for ib && jb {
		if i != j {
			return false
		}
		i, ib = r1()
		j, jb = r2()
	}
	if ib || jb {
		return false
	}
	return true
}

func inorder(node *TreeNode) func() (int, bool) {
	if node == nil {
		return func() (int, bool) {
			return 1, false
		}
	}
	var (
		res    int
		stack  []*TreeNode = []*TreeNode{node}
		stackB []bool      = []bool{false}
	)

	return func() (int, bool) {
		for len(stack) != 0 {
			curNode := stack[len(stack)-1]
			cur := stackB[len(stackB)-1]
			stack = stack[:len(stack)-1]
			stackB = stackB[:len(stackB)-1]
			if cur && curNode.Left == nil && curNode.Right == nil {
				res = curNode.Val
				return res, true
			} else if cur {
				continue
			} else {
				if curNode.Right != nil {
					stack = append(stack, curNode.Right)
					stackB = append(stackB, false)
				}
				stack = append(stack, curNode)
				stackB = append(stackB, true)
				if curNode.Left != nil {
					stack = append(stack, curNode.Left)
					stackB = append(stackB, false)
				}
			}
		}
		return 1, false
	}
}