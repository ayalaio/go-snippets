package main

import (
	"fmt"

	"github.com/ayalaio/go-snippets/algorithms/spanningtrees/kruskal"
)

func main() {

	vA := kruskal.NewVertex("A")
	vB := kruskal.NewVertex("B")
	vC := kruskal.NewVertex("C")
	vD := kruskal.NewVertex("D")
	vE := kruskal.NewVertex("E")
	vF := kruskal.NewVertex("F")
	vG := kruskal.NewVertex("G")

	edgeList := make([]*kruskal.Edge, 0)

	edgeList = append(edgeList, kruskal.NewEdge(vA, vB, 2))
	edgeList = append(edgeList, kruskal.NewEdge(vB, vE, 3))
	edgeList = append(edgeList, kruskal.NewEdge(vE, vD, 4))
	edgeList = append(edgeList, kruskal.NewEdge(vD, vG, 5))
	edgeList = append(edgeList, kruskal.NewEdge(vG, vF, 5))
	edgeList = append(edgeList, kruskal.NewEdge(vF, vA, 10))
	edgeList = append(edgeList, kruskal.NewEdge(vA, vC, 6))
	edgeList = append(edgeList, kruskal.NewEdge(vA, vE, 5))
	edgeList = append(edgeList, kruskal.NewEdge(vB, vD, 3))
	edgeList = append(edgeList, kruskal.NewEdge(vC, vD, 1))
	edgeList = append(edgeList, kruskal.NewEdge(vC, vF, 2))

	k := kruskal.NewKruskal(edgeList)

	keptEdgeList := k.GetKeptEdges()

	for _, e := range keptEdgeList {
		fmt.Printf("%s - %s : %d \n", e.PointA.Name, e.PointB.Name, e.Weight)
	}

}
