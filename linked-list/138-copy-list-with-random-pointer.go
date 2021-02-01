package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cur := head
	if head == nil {
		return nil
	}
	for cur != nil {
		node := &Node{Val: cur.Val, Next: nil}
		node.Next = cur.Next
		cur.Next = node
		cur = node.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	cur = head
	res := head.Next
	new_cur := res
	for cur != nil {
		cur.Next = cur.Next.Next
		if new_cur.Next != nil {
			new_cur.Next = new_cur.Next.Next
		}
		cur = cur.Next
		new_cur = new_cur.Next
	}
	return res
}

func copyRandomList2(head *Node) *Node {
	p := &Node{Val: 0, Next: nil, Random: nil}
	res := p
	cur := head
	m := make(map[*Node]*Node)
	for cur != nil {
		node := &Node{Val: cur.Val, Next: nil, Random: nil}
		p.Next = node
		m[cur] = node
		p = p.Next
		cur = cur.Next
	}
	cur = head
	p = res.Next
	for cur != nil {
		p.Random = m[cur.Random]
		p = p.Next
		cur = cur.Next
	}
	return res.Next
}

func main() {
	four := &Node{Val: 1, Next: nil}
	three := &Node{Val: 10, Next: four, Random: four}
	two := &Node{Val: 11, Next: three, Random: four}
	root := &Node{Val: 13, Next: two}
	copyRandomList(root)
}
