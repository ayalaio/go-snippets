package main

import (
	"container/heap"
	"fmt"
)

type Vertex struct {
	val       string
	neighbors []*Edge
}

func NewVertex(val string) *Vertex {
	return &Vertex{val, make([]*Edge, 0)}
}

type Edge struct {
	src, trg *Vertex
	weight   int
}

func NewEdge(src, trg *Vertex, weight int) *Edge {
	return &Edge{src, trg, weight}
}

func (self *Vertex) AddNeighbor(v *Vertex, weight int) {
	e := NewEdge(self, v, weight)
	self.neighbors = append(self.neighbors, e)
}

type EdgeHeap []*Edge

func (self EdgeHeap) Len() int           { return len(self) }
func (self EdgeHeap) Less(i, j int) bool { return self[i].weight < self[j].weight }
func (self EdgeHeap) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }
func (self *EdgeHeap) Push(x interface{}) {
	*self = append(*self, x.(*Edge))
}
func (self *EdgeHeap) Pop() interface{} {
	h := *self
	var el *Edge
	l := h.Len()
	el, *self = h[l-1], h[0:l-1]
	return el
}

type Dijkstra struct {
	precedent map[*Vertex]*Vertex
	weight    map[*Vertex]int
	edgePool  map[*Edge]bool
	edgeHeap  *EdgeHeap
	root      *Vertex
}

func NewDijkstra(root *Vertex) *Dijkstra {
	edgeHeap := &EdgeHeap{}
	heap.Init(edgeHeap)
	return &Dijkstra{
		make(map[*Vertex]*Vertex),
		make(map[*Vertex]int),
		make(map[*Edge]bool),
		edgeHeap,
		root}
}

func (self Dijkstra) Solve() {

	heap.Push(self.edgeHeap, NewEdge(self.root, self.root, 0))

	for self.edgeHeap.Len() > 0 {
		edge := heap.Pop(self.edgeHeap).(*Edge)
		trg := edge.trg
		src := edge.src

		currRegWeight, hasWeight := self.weight[trg]
		if !hasWeight || self.weight[src]+edge.weight < currRegWeight {
			self.weight[trg] = self.weight[src] + edge.weight
			self.precedent[trg] = src
		}

		for i := 0; i < len(trg.neighbors); i++ {
			nE := trg.neighbors[i]
			if _, isRegistered := self.edgePool[nE]; !isRegistered {
				heap.Push(self.edgeHeap, nE)
				self.edgePool[nE] = true
			}
		}

	}

}

func (self Dijkstra) Path(trg *Vertex) (path []string, weight int) {
	curr := trg
	path = []string{}
	for curr != nil && curr != self.root {
		path = append([]string{curr.val}, path...)
		curr = self.precedent[curr]
	}
	return append([]string{self.root.val}, path...), self.weight[trg]
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

	for i := 0; i < len(graph); i++ {
		d := NewDijkstra(graph[i])
		fmt.Println("Solving for ", graph[i].val)
		d.Solve()
		for j := 0; j < len(graph); j++ {
			s, w := d.Path(graph[j])
			fmt.Println(s, w)
		}
	}

}
