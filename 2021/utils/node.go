package utils

type Node struct {
	Content  interface{}
	Children []*Node
}

func NewNode(content interface{}) *Node {
	return &Node{Content: content, Children: make([]*Node, 0)}
}

func (self *Node) AddChildren(child *Node) {
	self.Children = append(self.Children, child)
}
