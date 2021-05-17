package prim

import (
	"container/heap"

	"github.com/daroay/go-snippets/algorithms/util"
	"github.com/daroay/utils/queue"
)

type Prim struct {
	root               *util.Vertex
	keptEdges          []*util.Edge
	isVertexVisitedMap map[*util.Vertex]bool
	edgeHeap           *util.EdgeHeap
}

func NewPrim(v *util.Vertex) *Prim {
	self := &Prim{
		v,
		make([]*util.Edge, 0),
		make(map[*util.Vertex]bool),
		&util.EdgeHeap{},
	}
	return self
}

func (self *Prim) GetKeptEdges() []*util.Edge {
	if len(self.keptEdges) == 0 {
		self.traveseVertex()
	}
	return self.keptEdges
}

func (self *Prim) isVertexVisited(v *util.Vertex) bool {
	_, ok := self.isVertexVisitedMap[v]
	return ok
}

func (self *Prim) visitVertex(v *util.Vertex) {
	self.isVertexVisitedMap[v] = true
}

func (self *Prim) traveseVertex() {

	l := queue.New()
	l.Add(self.root)

	for l.Len() > 0 {
		v := l.Front().Value.(*util.Vertex)
		l.Remove()
		self.visitVertex(v)

		for _, e := range v.EdgeList {
			heap.Push(self.edgeHeap, e)
		}

		expanded := false
		for !expanded && self.edgeHeap.Len() > 0 {
			ne := heap.Pop(self.edgeHeap).(*util.Edge)

			pA := ne.PointA
			pB := ne.PointB

			if !self.isVertexVisited(pA) {
				l.Add(pA)
				self.keptEdges = append(self.keptEdges, ne)
				expanded = true
			}

			if !self.isVertexVisited(pB) {
				l.Add(pB)
				self.keptEdges = append(self.keptEdges, ne)
				expanded = true
			}
		}

	}

}
