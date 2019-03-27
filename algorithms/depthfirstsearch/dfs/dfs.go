package dfs

import (
	"fmt"

	"github.com/ayalaio/utils/stack"
	"github.com/ayalaio/utils/vertex"
)

func DFS(l []*vertex.Vertex) {
	for _, v := range l {
		if !v.IsVisited {
			dfs(v)
		}
	}
}

func dfs(v *vertex.Vertex) {
	s := stack.New()
	s.Add(v)

	for s.Len() > 0 {
		v := s.Front().Value.(*vertex.Vertex)
		s.Remove()
		visit(v)
		for nve := v.Neighbors.Back(); nve != nil; nve = nve.Prev() {
			nv := nve.Value.(*vertex.Vertex)
			if !nv.IsVisited {
				s.Add(nv)
			}
		}
	}
}

func DFSRec(l []*vertex.Vertex) {
	for _, v := range l {
		if !v.IsVisited {
			dfsRec(v)
		}
	}
}

func dfsRec(v *vertex.Vertex) {
	visit(v)
	for nve := v.Neighbors.Front(); nve != nil; nve = nve.Next() {
		nv := nve.Value.(*vertex.Vertex)
		if !nv.IsVisited {
			dfsRec(nv)
		}
	}
}

func visit(v *vertex.Vertex) {
	v.Visit()
	fmt.Printf("Visited %d\n", v.Data)
}

type TopologicalOrder struct {
	order stack.Stack
}

func NewTopologicalOrder() *TopologicalOrder {
	t := TopologicalOrder{}
	t.order = stack.New()
	return &t
}

func (t *TopologicalOrder) GetOrder(l []*vertex.Vertex) {
	for _, v := range l {
		if !v.IsVisited {

		}
	}
}
