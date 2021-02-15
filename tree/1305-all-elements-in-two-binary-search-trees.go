package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res := []int{}
	one := inorder(root1)
	two := inorder(root2)
	i, ib := one()
	j, jb := two()
	for ib && jb {
		if i <= j {
			res = append(res, i)
			i, ib = one()
		} else {
			res = append(res, j)
			j, jb = two()
		}
	}
	if ib {
		res = append(res, i)
	} else {
		res = append(res, j)
	}
	for ib {
		i, ib = one()
		if ib {
			res = append(res, i)
		}

	}
	for jb {
		j, jb = two()
		if jb {
			res = append(res, j)
		}
	}
	return res
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
			if cur {
				res = curNode.Val
				return res, true
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
