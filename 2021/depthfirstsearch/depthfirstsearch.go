package depthfirstsearch

import "github.com/daroay/go-snippets/2021/utils"

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
