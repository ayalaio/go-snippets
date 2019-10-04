package main

import (
	"container/heap"
	"fmt"
)

type Vertex struct {
	Label    string
	Children []*Edge
}

func NewVertex(label string) *Vertex {
	return &Vertex{label, make([]*Edge, 0)}
}

func (self *Vertex) AddNeighbor(v *Vertex, w int) {
	e := NewEdge(w, self, v)
	self.Children = append(self.Children, e)
	v.Children = append(v.Children, e)
}

type Edge struct {
	Weight   int
	Src, Dst *Vertex
}

func NewEdge(w int, src, dst *Vertex) *Edge {
	return &Edge{w, src, dst}
}

type EdgeHeap []*Edge

func (self EdgeHeap) Len() int           { return len(self) }
func (self EdgeHeap) Less(i, j int) bool { return self[i].Weight < self[j].Weight }
func (self EdgeHeap) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (self *EdgeHeap) Pop() interface{} {
	h := *self
	edge := h[len(h)-1]
	*self = h[0 : len(h)-1]
	return edge
}

func (self *EdgeHeap) Push(x interface{}) {
	*self = append(*self, x.(*Edge))
}

func RunDijkstra(root *Vertex, graph []*Vertex) (map[*Vertex]*Vertex, map[*Vertex]int) {
	precedents := make(map[*Vertex]*Vertex)
	weights := make(map[*Vertex]int)
	isEdgeAdded := make(map[*Edge]bool)
	edgeHeap := &EdgeHeap{}
	heap.Push(edgeHeap, NewEdge(0, root, root))

	for edgeHeap.Len() > 0 {
		var edge *Edge
		edge = heap.Pop(edgeHeap).(*Edge)

		weightThroughThisPath := weights[edge.Src] + edge.Weight

		val, hasWeight := weights[edge.Dst]

		if !hasWeight || val > weightThroughThisPath {
			weights[edge.Dst] = weightThroughThisPath
			precedents[edge.Dst] = edge.Src
		}

		for _, child := range edge.Dst.Children {
			if _, ok := isEdgeAdded[child]; !ok {
				isEdgeAdded[child] = true
				heap.Push(edgeHeap, child)
			}
		}

	}

	return precedents, weights
}

func main() {

	a := NewVertex("A")
	b := NewVertex("B")
	c := NewVertex("C")
	d := NewVertex("D")
	e := NewVertex("E")
	f := NewVertex("F")
	g := NewVertex("G")

	a.AddNeighbor(b, 4)
	a.AddNeighbor(c, 3)
	a.AddNeighbor(e, 7)

	b.AddNeighbor(a, 4)
	b.AddNeighbor(d, 5)
	b.AddNeighbor(c, 6)

	c.AddNeighbor(a, 3)
	c.AddNeighbor(b, 6)
	c.AddNeighbor(d, 11)
	c.AddNeighbor(e, 8)

	d.AddNeighbor(b, 5)
	d.AddNeighbor(c, 11)
	d.AddNeighbor(e, 2)
	d.AddNeighbor(g, 10)
	d.AddNeighbor(f, 2)

	e.AddNeighbor(a, 7)
	e.AddNeighbor(c, 8)
	e.AddNeighbor(d, 2)
	e.AddNeighbor(g, 5)

	f.AddNeighbor(d, 2)
	f.AddNeighbor(g, 3)

	g.AddNeighbor(e, 5)
	g.AddNeighbor(d, 10)
	g.AddNeighbor(f, 3)

	graph := []*Vertex{a, b, c, d, e, f, g}

	_ = graph

	precedents, weights := RunDijkstra(a, graph)

	for to, from := range precedents {
		fmt.Printf("from %s to %s is %d\n", from.Label, to.Label, weights[to])
	}

}
