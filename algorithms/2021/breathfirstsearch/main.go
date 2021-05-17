package main

import (
	"fmt"

	"github.com/daroay/go-snippets/algorithms/2021/utils"
)

func BreathfirstSearch(node *utils.Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Number)
}

func main() {
	fmt.Println("Breathfirst search")
	root := utils.NewNode(0)
	one := utils.NewNode(1)
	two := utils.NewNode(2)
	three := utils.NewNode(3)
	four := utils.NewNode(4)
	five := utils.NewNode(5)
	six := utils.NewNode(6)
	root.AddChildren(one)
	root.AddChildren(two)
	one.AddChildren(three)
	one.AddChildren(four)
	two.AddChildren(five)
	two.AddChildren(six)
	BreathfirstSearch(root)

	s := utils.Stack{}
	s.Push(1)
	s.Push(2)

	fmt.Println(s)

	for len(s) > 0 {
		fmt.Println(s.Pop())
	}
}
