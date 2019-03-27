package topologicalorder

import (
	"github.com/ayalaio/utils/stack"
	"github.com/ayalaio/utils/vertex"
)

type TopologicalOrder struct {
	order stack.Stack
}

func New() *TopologicalOrder {
	t := TopologicalOrder{}
	t.order = stack.New()
	return &t
}

func (t *TopologicalOrder) dfs(v *vertex.Vertex) {
	v.Visit()
	for ne := v.Neighbors.Front(); ne != nil; ne = ne.Next() {
		nve := ne.Value.(*vertex.Vertex)
		if !nve.IsVisited {
			t.dfs(nve)
		}
	}
	t.order.Add(v)
}

func (t *TopologicalOrder) GetOrder(l []*vertex.Vertex) stack.Stack {
	for _, v := range l {
		if !v.IsVisited {
			t.dfs(v)
		}
	}
	return t.order
}
