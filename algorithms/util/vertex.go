package util

type Vertex struct {
	Name     string
	EdgeList []*Edge
}

func NewVertex(name string) *Vertex {
	return &Vertex{name, make([]*Edge, 0)}
}

func (self *Vertex) AddEdge(e *Edge) {
	self.EdgeList = append(self.EdgeList, e)
}
