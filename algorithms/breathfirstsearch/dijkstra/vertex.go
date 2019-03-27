package dijkstra

import (
	"container/list"
)

// Vertex is the vertex of a Graph
type Vertex struct {
	Data      interface{}
	IsVisited bool
	Neighbors *list.List //<Edge>
}

// NewVertex gives a new Vertex
func NewVertex(data string) (v *Vertex) {
	v = &Vertex{}
	v.Data = data
	v.IsVisited = false
	v.Neighbors = list.New()
	return
}

// AddNeighbor adds a new neighbor Vertex to the Vertex
func (v *Vertex) AddNeighbor(n *Vertex, weight int) {
	edge := Edge{n, weight}
	v.Neighbors.PushBack(edge)
}

// Visit marks the Vertex as visited
func (v *Vertex) Visit() {
	v.IsVisited = true
}
