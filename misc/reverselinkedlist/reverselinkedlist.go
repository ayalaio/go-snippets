package main

type Node struct {
	Content int
	Next    *Node
}

func reverse(rootPtr **Node) {
	root := *rootPtr

	curr := root
	var prev *Node
	var fwd *Node

	for curr != nil {
		fwd = curr.Next
		curr.Next = prev
		prev = curr
		curr = fwd
	}

	*rootPtr = prev

}

func main() {
	root := &Node{0, nil}

	curr := root
	for i := 1; i < 10; i++ {
		curr.Next = &Node{i, nil}
		curr = curr.Next
	}

	reverse(&root)

	for x := root; x != nil; x = x.Next {
		println(x.Content)
	}

}
