package kruskal

type Edge struct {
	PointA *Vertex
	PointB *Vertex
	Weight int
}

func NewEdge(u, v *Vertex, w int) *Edge {
	return &Edge{u, v, w}
}
