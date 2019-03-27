package bfs

import (
	"fmt"

	"github.com/ayalaio/utils/queue"
	"github.com/ayalaio/utils/vertex"
)

// BFS Visits nodes
func BFS(root *vertex.Vertex) {
	q := queue.New()
	q.Add(root)

	for q.Len() > 0 {
		e := q.Front()
		v := e.Value.(*vertex.Vertex)
		visit(v)
		for neighbor := v.Neighbors.Front(); neighbor != nil; neighbor = neighbor.Next() {
			neighborVertex := neighbor.Value.(*vertex.Vertex)
			if !neighborVertex.IsVisited {
				q.Add(neighborVertex)
			}
		}
		q.Remove()
	}
}

func visit(v *vertex.Vertex) {
	v.Visit()
	fmt.Printf("Visited %d\n", v.Data)
}
