package main

import (
	"fmt"

	"github.com/daroay/go-snippets/algorithms/util"
)

func BreathfirstSearch(node *util.Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Number)
}

func main() {
	fmt.Println("Breathfirst search")
	root := util.NewNode(0)
	one := util.NewNode(1)
	two := util.NewNode(2)
	three := util.NewNode(3)
	four := util.NewNode(4)
	five := util.NewNode(5)
	six := util.NewNode(6)
	root.AddChildren(one)
	root.AddChildren(two)
	one.AddChildren(three)
	one.AddChildren(four)
	two.AddChildren(five)
	two.AddChildren(six)
	BreathfirstSearch(root)

	s := util.Stack{}
	s.Push(1)
	s.Push(2)

	fmt.Println(s)

	for len(s) > 0 {
		fmt.Println(s.Pop())
	}
}
