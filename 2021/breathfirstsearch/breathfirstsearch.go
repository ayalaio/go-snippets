package breathfirstsearch

import (
	"github.com/daroay/go-snippets/2021/utils"
)

func AddChildrenToQueue(children []*utils.Node, q *utils.Queue) {
	for _, c := range children {
		q.Push(c)
	}
}

func BreathFirstSearch(root *utils.Node, block func(*utils.Node)) {
	if root == nil {
		return
	}
	q := utils.Queue{}
	q.Push(root)
	for item := q.Take(); item != nil; item = q.Take() {
		n := item.(*utils.Node)
		block(n)
		AddChildrenToQueue(n.Children, &q)
	}
}
