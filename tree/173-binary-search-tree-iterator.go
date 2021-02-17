package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
	visit []bool
	f     func() int
}

func Constructor(root *TreeNode) BSTIterator {
	b := &BSTIterator{
		stack: []*TreeNode{root},
		visit: []bool{false},
	}
	f := func() int {
		for len(b.stack) > 0 {
			node := b.stack[len(b.stack)-1]
			visited := b.visit[len(b.visit)-1]
			b.stack = b.stack[:len(b.stack)-1]
			b.visit = b.visit[:len(b.visit)-1]
			if visited {
				return node.Val
			} else {
				if node.Right != nil {
					b.stack = append(b.stack, node.Right)
					b.visit = append(b.visit, false)
				}
				b.stack = append(b.stack, node)
				b.visit = append(b.visit, true)
				if node.Left != nil {
					b.stack = append(b.stack, node.Left)
					b.visit = append(b.visit, false)
				}
			}
		}
		return 0
	}
	b.f = f
	return *b
}

func (this *BSTIterator) Next() int {
	x := len(this.stack)
	fmt.Println(x)
	return this.f()
}

func (this *BSTIterator) HasNext() bool {
	x := len(this.stack)
	return x > 0
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
