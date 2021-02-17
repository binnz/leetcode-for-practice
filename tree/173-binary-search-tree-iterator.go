package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
	visit []bool
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{
		stack: []*TreeNode{root},
		visit: []bool{false},
	}
	return b
}

func (this *BSTIterator) f() int {
	for len(this.stack) > 0 {
		node := this.stack[len(this.stack)-1]
		visited := this.visit[len(this.visit)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.visit = this.visit[:len(this.visit)-1]
		if visited {
			return node.Val
		} else {
			if node.Right != nil {
				this.stack = append(this.stack, node.Right)
				this.visit = append(this.visit, false)
			}
			this.stack = append(this.stack, node)
			this.visit = append(this.visit, true)
			if node.Left != nil {
				this.stack = append(this.stack, node.Left)
				this.visit = append(this.visit, false)
			}
		}
	}
	return 0
}

func (this *BSTIterator) Next() int {
	return this.f()
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func main() {
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 2}
	root := &TreeNode{Val: 3, Left: l, Right: r}
	target := Constructor(root)
	target.Next()
	target.HasNext()
	target.Next()
	target.HasNext()
}
