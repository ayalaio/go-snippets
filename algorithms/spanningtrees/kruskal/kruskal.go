package kruskal

import (
	"container/heap"
	"fmt"
)

type EdgeHeap []*Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Kruskal struct {
	edgeHeap  *EdgeHeap
	djsM      map[*Vertex]*DisjointSet
	keptEdges []*Edge
}

func NewKruskal(eList []*Edge) *Kruskal {
	self := &Kruskal{
		&EdgeHeap{},
		make(map[*Vertex]*DisjointSet),
		make([]*Edge, 0),
	}
	heap.Init(self.edgeHeap)
	for _, e := range eList {
		self.pushEdge(e)
	}
	return self
}

func (self *Kruskal) GetKeptEdges() []*Edge {
	if len(self.keptEdges) == 0 {
		self.traverseEdges()
	}
	return self.keptEdges
}

func (self *Kruskal) keepEdge(e *Edge) {
	self.keptEdges = append(self.keptEdges, e)
}

func (self *Kruskal) pushEdge(e *Edge) {
	heap.Push(self.edgeHeap, e)
}

func (self *Kruskal) popEdge() *Edge {
	return heap.Pop(self.edgeHeap).(*Edge)
}

func (self *Kruskal) GetDisjointSet(v *Vertex) *DisjointSet {
	var djs *DisjointSet

	if val, ok := self.djsM[v]; ok {
		djs = val
	} else {
		djs = NewDisjointSet(v, nil)
		self.djsM[v] = djs
	}

	return djs
}

func (self *Kruskal) traverseEdges() {

	for len(*self.edgeHeap) > 0 {
		e := heap.Pop(self.edgeHeap).(*Edge)
		dA := self.GetDisjointSet(e.PointA)
		dB := self.GetDisjointSet(e.PointB)

		fmt.Println(e.Weight)

		err := MergeDisjointSet(dA, dB)
		if err == nil {
			self.keepEdge(e)
		}
	}

}
