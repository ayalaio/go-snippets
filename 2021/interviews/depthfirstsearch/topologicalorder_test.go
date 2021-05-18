package depthfirstsearch

import (
	"fmt"
	"testing"

	"github.com/daroay/go-snippets/2021/utils"
)

var isBeingVisited map[*utils.Node]bool = make(map[*utils.Node]bool)
var isAddedToSolution map[*utils.Node]bool = make(map[*utils.Node]bool)

func TopologicalOrder(n *utils.Node) ([]*utils.Node, error) {
	if isBeingVisited[n] {
		return nil, fmt.Errorf("Graph has cycles")
	}
	if isAddedToSolution[n] {
		return []*utils.Node{}, nil
	}
	isBeingVisited[n] = true
	isAddedToSolution[n] = true

	list := []*utils.Node{n}
	for _, c := range n.Children {
		orderForChild, err := TopologicalOrder(c)
		if err != nil {
			return nil, err
		}
		list = append(list, orderForChild...)
	}
	isBeingVisited[n] = false
	return list, nil
}

func Test_TopologicalOrder_WithCycle_ShouldFail(t *testing.T) {
	a := utils.NewNode("A")
	b := utils.NewNode("B")
	c := utils.NewNode("C")
	d := utils.NewNode("D")
	e := utils.NewNode("E")
	f := utils.NewNode("F")

	root := d
	d.AddChildren(a)
	a.AddChildren(c)
	c.AddChildren(b)
	c.AddChildren(f)
	a.AddChildren(e)
	e.AddChildren(d) // cycle

	_, err := TopologicalOrder(root)
	if err == nil {
		t.Errorf("Error was expected")
	}
}

func Test_TopologicalOrder_WithoutCycle_ShouldSucceed(t *testing.T) {
	a := utils.NewNode("A")
	b := utils.NewNode("B")
	c := utils.NewNode("C")
	d := utils.NewNode("D")
	e := utils.NewNode("E")
	f := utils.NewNode("F")

	root := d
	d.AddChildren(a)
	a.AddChildren(c)
	c.AddChildren(b)
	c.AddChildren(f)
	a.AddChildren(e)
	e.AddChildren(c)

	order, err := TopologicalOrder(root)
	if err != nil {
		t.Errorf(err.Error())
	}

	expectedOrder := []string{"D", "A", "C", "B", "F", "E"}

	for i := 0; i < len(order); i++ {
		if expectedOrder[i] != order[i].Content {
			t.Errorf("Order is wrong %s should be %s", order[i].Content, expectedOrder[i])
		}
	}
}
