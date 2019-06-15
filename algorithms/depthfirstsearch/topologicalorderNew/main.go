package main

import "fmt"

type Node struct {
	value    string
	children []*Node
}

type TopologicalOrder struct {
	visitMap, verified map[*Node]bool
	compiledList       []string
}

func (self *Node) addChildren(child *Node) {
	self.children = append(self.children, child)
}

func NewNode(val string) *Node {
	return &Node{val, make([]*Node, 0)}
}

func FindTopologicalOrder(nodes []*Node) ([]string, error) {

	to := &TopologicalOrder{
		visitMap:     make(map[*Node]bool),
		verified:     make(map[*Node]bool),
		compiledList: make([]string, 0),
	}

	var err error

	for i := 0; i < len(nodes); i++ {
		currNode := nodes[i]
		if _, ok := to.verified[currNode]; !ok {
			err = to.findTopologicalOrder(currNode)
		}
		if err != nil {
			return nil, err
		}
	}

	return to.compiledList, err

}

func (self *TopologicalOrder) findTopologicalOrder(node *Node) error {

	if _, ok := self.visitMap[node]; ok {
		return fmt.Errorf("Graph has cycles")
	}

	if _, ok := self.verified[node]; ok {
		return nil
	}

	self.visitMap[node] = true
	for i := 0; i < len(node.children); i++ {
		err := self.findTopologicalOrder(node.children[i])
		if err != nil {
			return err
		}
	}
	self.compiledList = append([]string{node.value}, self.compiledList...)
	delete(self.visitMap, node)
	self.verified[node] = true
	return nil
}

func main() {

	var a, b, c, d, e, f *Node
	var order []string
	var nodes []*Node
	var err error

	a = NewNode("A")
	b = NewNode("B")
	c = NewNode("C")
	d = NewNode("D")
	e = NewNode("E")
	f = NewNode("F")

	a.addChildren(b)
	a.addChildren(c)

	b.addChildren(c)

	c.addChildren(d)

	d.addChildren(e)
	d.addChildren(f)

	e.addChildren(f)

	nodes = []*Node{b, c, a, d, f, e}

	order, err = FindTopologicalOrder(nodes)

	fmt.Println(order, err)

	a = NewNode("A")
	b = NewNode("B")
	c = NewNode("C")
	d = NewNode("D")
	e = NewNode("E")
	f = NewNode("F")

	a.addChildren(b)
	a.addChildren(c)

	b.addChildren(c)

	c.addChildren(a)

	d.addChildren(e)
	d.addChildren(f)

	e.addChildren(f)

	nodes = []*Node{b, c, a, d, f, e}

	order, err = FindTopologicalOrder(nodes)

	fmt.Println(order, err)

}
