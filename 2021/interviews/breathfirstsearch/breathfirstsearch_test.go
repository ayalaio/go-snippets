package breathfirstsearch

import (
	"testing"

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

func TestBreathFirstSearch(t *testing.T) {
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
	currentNumber := 0
	BreathFirstSearch(root, func(n *utils.Node) {
		if n.Number != currentNumber {
			t.Errorf("Expecting %d, got %d", currentNumber, n.Number)
		}
		currentNumber++
	})
}
