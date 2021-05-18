package depthfirstsearch

import (
	"testing"

	"github.com/daroay/go-snippets/2021/utils"
)

func AddChildrenToStack(children []*utils.Node, s *utils.Stack) {
	for i := len(children) - 1; i >= 0; i-- {
		s.Push(children[i])
	}
}

func DepthFirstSearch(node *utils.Node, block func(*utils.Node)) {
	s := utils.Stack{}
	s.Push(node)
	for item := s.Pop(); item != nil; item = s.Pop() {
		n := item.(*utils.Node)
		block(n)
		AddChildrenToStack(n.Children, &s)
	}
}

func TestDepthFirstSearch(t *testing.T) {
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
	expectedOrder := []int{0, 1, 3, 4, 2, 5, 6}
	expectedOrderIter := 0
	DepthFirstSearch(root, func(n *utils.Node) {
		expectedNumber := expectedOrder[expectedOrderIter]
		if n.Number != expectedNumber {
			t.Errorf("Expected %d, got %d", expectedNumber, n.Number)
		}
		expectedOrderIter++
	})
}
