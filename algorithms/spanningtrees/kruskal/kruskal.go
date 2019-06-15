package kruskal

import (
	"container/heap"

	"github.com/ayalaio/go-snippets/algorithms/util"
)

type Kruskal struct {
	edgeHeap  *util.EdgeHeap
	djsM      map[*util.Vertex]*DisjointSet
	keptEdges []*util.Edge
}

func NewKruskal(eList []*util.Edge) *Kruskal {
	self := &Kruskal{
		&util.EdgeHeap{},
		make(map[*util.Vertex]*DisjointSet),
		make([]*util.Edge, 0),
	}
	heap.Init(self.edgeHeap)
	for _, e := range eList {
		self.pushEdge(e)
	}
	return self
}

func (self *Kruskal) GetKeptEdges() []*util.Edge {
	if len(self.keptEdges) == 0 {
		self.traverseEdges()
	}
	return self.keptEdges
}

func (self *Kruskal) keepEdge(e *util.Edge) {
	self.keptEdges = append(self.keptEdges, e)
}

func (self *Kruskal) pushEdge(e *util.Edge) {
	heap.Push(self.edgeHeap, e)
}

func (self *Kruskal) popEdge() *util.Edge {
	return heap.Pop(self.edgeHeap).(*util.Edge)
}

func (self *Kruskal) GetDisjointSet(v *util.Vertex) *DisjointSet {
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

	for self.edgeHeap.Len() > 0 {
		e := self.popEdge()
		dA := self.GetDisjointSet(e.PointA)
		dB := self.GetDisjointSet(e.PointB)

		err := MergeDisjointSet(dA, dB)
		if err == nil {
			self.keepEdge(e)
		}
	}

}
