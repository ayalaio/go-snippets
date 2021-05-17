package utils

type Node struct {
	Number   int
	Children []*Node
}

func NewNode(num int) *Node {
	return &Node{Number: num, Children: make([]*Node, 0)}
}

func (self *Node) AddChildren(child *Node) {
	self.Children = append(self.Children, child)
}
