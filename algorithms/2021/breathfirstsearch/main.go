package main

import (
	"fmt"

	node "github.com/daroay/2021/utils/Node.go"
	stack "github.com/daroay/2021/utils/Stack.go"
)

func BreathfirstSearch(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Number)
}

func main() {
	fmt.Println("Breathfirst search")
	root := node.NewNode(0)
	one := node.NewNode(1)
	two := node.NewNode(2)
	three := node.NewNode(3)
	four := node.NewNode(4)
	five := node.NewNode(5)
	six := node.NewNode(6)
	root.AddChildren(one)
	root.AddChildren(two)
	one.AddChildren(three)
	one.AddChildren(four)
	two.AddChildren(five)
	two.AddChildren(six)
	BreathfirstSearch(root)

	s := stack.Stack{}
	s.Push(1)
	s.Push(2)

	fmt.Println(s)

	for len(s) > 0 {
		fmt.Println(s.Pop())
	}
}
