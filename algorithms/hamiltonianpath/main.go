package main

import (
	"github.com/ayalaio/go-snippets/algorithms/hamiltonianpath/hamilton"
	"github.com/ayalaio/go-snippets/algorithms/util"
)

func main() {

	println("------ kruskal --------")

	vA := util.NewVertex("A")
	vB := util.NewVertex("B")
	vC := util.NewVertex("C")
	vD := util.NewVertex("D")
	vE := util.NewVertex("E")
	vF := util.NewVertex("F")
	vG := util.NewVertex("G")

	vertexList := make([]*util.Vertex, 0)

	vertexList = append(vertexList, vA)
	vertexList = append(vertexList, vB)
	vertexList = append(vertexList, vC)
	vertexList = append(vertexList, vD)
	vertexList = append(vertexList, vE)
	vertexList = append(vertexList, vF)
	vertexList = append(vertexList, vG)

	edgeList := make([]*util.Edge, 0)

	edgeList = append(edgeList, util.NewEdge(vA, vB, 2))
	edgeList = append(edgeList, util.NewEdge(vB, vE, 3))
	edgeList = append(edgeList, util.NewEdge(vE, vD, 4))
	edgeList = append(edgeList, util.NewEdge(vD, vG, 5))
	edgeList = append(edgeList, util.NewEdge(vG, vF, 5))
	edgeList = append(edgeList, util.NewEdge(vF, vA, 10))
	edgeList = append(edgeList, util.NewEdge(vA, vC, 6))
	edgeList = append(edgeList, util.NewEdge(vA, vE, 5))
	edgeList = append(edgeList, util.NewEdge(vB, vD, 3))
	edgeList = append(edgeList, util.NewEdge(vC, vD, 1))
	edgeList = append(edgeList, util.NewEdge(vC, vF, 2))

	println("-------- hamiltonian / not a spanning tree -------")

	h := hamilton.New(vertexList)
	h.Solve()

}
