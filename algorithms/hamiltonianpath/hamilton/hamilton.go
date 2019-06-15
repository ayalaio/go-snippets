package hamilton

import (
	"fmt"

	"github.com/ayalaio/go-snippets/algorithms/util"
)

type HamiltonianPath struct {
	visitMap         map[*util.Vertex]int
	totalNumOfVertex int
	vertexList       []*util.Vertex
	selectedEdgeMap  map[*util.Edge]int
}

func New(vertexList []*util.Vertex) *HamiltonianPath {
	return &HamiltonianPath{
		visitMap:         make(map[*util.Vertex]int, 0),
		totalNumOfVertex: len(vertexList),
		vertexList:       vertexList,
		selectedEdgeMap:  make(map[*util.Edge]int, 0),
	}
}

func (self *HamiltonianPath) Print() {
	for k, isSelected := range self.selectedEdgeMap {
		if isSelected > 0 {
			fmt.Println(k.PointA.Name, " - ", k.PointB.Name)
		}
	}
}

func (self *HamiltonianPath) Solve() {
	for i, v := range self.vertexList {
		fmt.Printf("solving with %d\n", i)
		solved := self.solve(v)
		if solved {
			self.Print()
		}
		self = New(self.vertexList)
	}
}

func (self *HamiltonianPath) solve(v *util.Vertex) bool {

	if self.isVisited(v) {
		return false
	}

	self.visit(v)

	if self.isAllVisited() {
		return true
	}

	for _, edge := range v.EdgeList {
		other, _ := edge.ContraPart(v)

		self.selectEdge(edge)

		solved := self.solve(other)
		if solved {
			return true
		}
		self.unselectEdge(edge)

	}
	self.unvisit(v)
	return false

}

func (self *HamiltonianPath) isVisited(v *util.Vertex) bool {
	val, ok := self.visitMap[v]
	if val > 0 && ok {
		return true
	}
	return false
}

func (self *HamiltonianPath) visit(v *util.Vertex) {
	self.visitMap[v]++
}

func (self *HamiltonianPath) unvisit(v *util.Vertex) {
	self.visitMap[v]--
}

func (self *HamiltonianPath) selectEdge(e *util.Edge) {
	self.selectedEdgeMap[e]++
}

func (self *HamiltonianPath) unselectEdge(e *util.Edge) {
	self.selectedEdgeMap[e]--
}

func (self *HamiltonianPath) isAllVisited() bool {
	currentTotal := 0
	for _, isVisited := range self.visitMap {
		if isVisited > 0 {
			currentTotal++
		}
	}
	return currentTotal == self.totalNumOfVertex
}
