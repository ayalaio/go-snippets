package util

import (
	"fmt"
)

type Edge struct {
	PointA *Vertex
	PointB *Vertex
	Weight int
}

func NewEdge(u, v *Vertex, w int) *Edge {
	self := &Edge{u, v, w}
	u.AddEdge(self)
	v.AddEdge(self)
	return self
}

func (self *Edge) ContraPart(v *Vertex) (*Vertex, error) {
	var u *Vertex
	if self.PointB == v {
		u = self.PointA
	} else if self.PointA == v {
		u = self.PointB
	} else {
		return nil, fmt.Errorf("Wrong parameter")
	}
	return u, nil
}
