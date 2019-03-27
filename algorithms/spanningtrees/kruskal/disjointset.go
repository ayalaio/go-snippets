package kruskal

import "fmt"

type DisjointSet struct {
	v      *Vertex
	parent *DisjointSet
	rank   int
}

func NewDisjointSet(v *Vertex, parent *DisjointSet) *DisjointSet {
	self := &DisjointSet{v, parent, 0}
	if parent == nil {
		self.parent = self
	}
	return self
}

func (self *DisjointSet) Find() *DisjointSet {
	var djs *DisjointSet
	if self.parent == self {
		djs = self
	} else {
		self.parent = self.parent.Find()
		self.rank = 1
		djs = self.parent
	}
	return djs
}

func MergeDisjointSet(u, v *DisjointSet) error {

	pU := u.Find()
	pV := v.Find()

	if pU != pV {
		if pU.rank == pV.rank {
			pV.parent = pU
			pU.rank++
		} else if pU.rank > pV.rank {
			pV.parent = pU
		} else if pU.rank < pV.rank {
			pU.parent = pV
		}

		return nil

	} else {
		return fmt.Errorf("Sets are already merged")
	}

}
