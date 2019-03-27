package kruskal

type Vertex struct {
	Name    string
	Visited bool
}

func NewVertex(name string) *Vertex {
	return &Vertex{name, false}
}
