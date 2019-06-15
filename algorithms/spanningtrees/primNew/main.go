package main

import (
	"container/heap"
	"fmt"
)

type EdgeHeap []*Edge

func (self EdgeHeap) Len() int           { return len(self) }
func (self EdgeHeap) Less(i, j int) bool { return self[i].weight < self[j].weight }
func (self EdgeHeap) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (self *EdgeHeap) Push(x interface{}) {
	*self = append(*self, x.(*Edge))
}

func (self *EdgeHeap) Pop() interface{} {
	h := *self
	var el interface{}
	l := len(h)
	el, *self = h[l-1], h[0:l-1]
	return el
}

type Prim struct {
	edgeCandidateHeap *EdgeHeap
	visitedVertex     map[*Vertex]bool
	reachableVertex   map[*Vertex]bool
	selectedEdges     map[*Edge]bool
	totalNumOfVertex  int
}

func (self Prim) isSolved() bool {
	return len(self.reachableVertex) == self.totalNumOfVertex
}

func NewPrim(totalNumOfVertex int) *Prim {
	h := &EdgeHeap{}
	heap.Init(h)
	return &Prim{
		h,
		make(map[*Vertex]bool),
		make(map[*Vertex]bool),
		make(map[*Edge]bool),
		totalNumOfVertex,
	}
}

type Vertex struct {
	val       string
	neighbors []*Edge
}

func (self *Vertex) String() string {
	return fmt.Sprint(self.val)
}

func NewVertex(value string) *Vertex {
	return &Vertex{value, make([]*Edge, 0)}
}

type Edge struct {
	u, v   *Vertex
	weight int
}

func (self Edge) String() string {
	return fmt.Sprint("(", self.u.val, ",", self.v.val, ")@", self.weight)
}

func NewEdge(u, v *Vertex, weight int) *Edge {
	return &Edge{u, v, weight}
}

func (self *Vertex) AddNeighbor(v *Vertex, weight int) {
	e := NewEdge(self, v, weight)
	self.neighbors = append(self.neighbors, e)
	v.neighbors = append(v.neighbors, e)
}

func (self *Prim) spanningTree(node *Vertex) {

	if _, ok := self.visitedVertex[node]; ok {
		return
	}

	self.visitedVertex[node] = true

	// each node's edge NOT YET selected add to the candidate pool
	for i := 0; i < len(node.neighbors); i++ {
		if _, ok := self.selectedEdges[node.neighbors[i]]; !ok {
			heap.Push(self.edgeCandidateHeap, node.neighbors[i])
		}
	}

	// pick the min edge from the candidates
	currEdge := heap.Pop(self.edgeCandidateHeap).(*Edge)

	// check if the edge's nodes are already reachable
	_, uAlreadyReachable := self.reachableVertex[currEdge.u]
	_, vAlreadyReachable := self.reachableVertex[currEdge.v]

	// if u is not reachable yet, reach it and select the edge
	if !uAlreadyReachable {
		self.reachableVertex[currEdge.u] = true
		self.selectedEdges[currEdge] = true
	}

	// if v is not reachable yet, reach it and select the edge
	if !vAlreadyReachable {
		self.reachableVertex[currEdge.v] = true
		self.selectedEdges[currEdge] = true
	}

	// Try to visit the edge's nodes,
	// they'll be rejected if they are already expanded
	self.spanningTree(currEdge.u)
	self.spanningTree(currEdge.v)

	return
}

func SpanningTree(nodes []*Vertex) {
	for i := 0; i < len(nodes); i++ {
		fmt.Println("--- Starting new from ", nodes[i].val, " ---")
		prim := NewPrim(len(nodes))
		prim.spanningTree(nodes[i])
		if prim.isSolved() {
			fmt.Println(prim.selectedEdges)
		}
	}
}

func main() {
	a := NewVertex("A")
	b := NewVertex("B")
	c := NewVertex("C")
	d := NewVertex("D")
	e := NewVertex("E")
	f := NewVertex("F")
	g := NewVertex("G")

	vertexList := []*Vertex{a, b, c, d, e, f, g}

	a.AddNeighbor(b, 2)
	b.AddNeighbor(e, 3)
	e.AddNeighbor(d, 4)
	d.AddNeighbor(g, 5)
	g.AddNeighbor(d, 5)
	g.AddNeighbor(f, 3)
	f.AddNeighbor(a, 10)
	a.AddNeighbor(c, 6)
	a.AddNeighbor(e, 5)
	b.AddNeighbor(d, 3)
	c.AddNeighbor(d, 1)
	c.AddNeighbor(f, 2)

	SpanningTree(vertexList)

}
