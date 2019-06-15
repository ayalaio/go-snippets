package main

type Node struct {
	Content int
	Next    *Node
}

func findMiddle(n *Node) *Node {
	var l1 *Node
	var l2 *Node

	l1, l2 = n, n

	for l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next
		if l2 != nil {
			l2 = l2.Next
		} else {
			break
		}
	}

	return l1
}

func main() {
	root := &Node{0, nil}

	curr := root
	for i := 1; i < 3; i++ {
		curr.Next = &Node{i, nil}
		curr = curr.Next
	}

	m := findMiddle(root)
	println(m.Content)
}
