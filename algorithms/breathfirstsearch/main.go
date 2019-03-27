package main

import (
	"container/list"

	"github.com/ayalaio/go-snippets/algorithms/breathfirstsearch/bfs"
	"github.com/ayalaio/go-snippets/algorithms/breathfirstsearch/dijkstra"
	"github.com/ayalaio/utils/vertex"
)

func main() {

	println("--------- bfs ------------")

	v1 := vertex.NewVertex(1)
	v2 := vertex.NewVertex(2)
	v3 := vertex.NewVertex(3)
	v4 := vertex.NewVertex(4)
	v5 := vertex.NewVertex(5)

	v1.AddNeighbor(v2)
	v1.AddNeighbor(v4)
	v2.AddNeighbor(v3)
	v4.AddNeighbor(v5)

	bfs.BFS(v1)

	println("--------- dijkstra ------------")

	dA := dijkstra.NewVertex("A")
	dB := dijkstra.NewVertex("B")
	dC := dijkstra.NewVertex("C")
	dD := dijkstra.NewVertex("D")
	dE := dijkstra.NewVertex("E")
	dF := dijkstra.NewVertex("F")
	dG := dijkstra.NewVertex("G")
	dH := dijkstra.NewVertex("H")

	myGraph := list.New()
	myGraph.PushFront(dA)
	myGraph.PushFront(dB)
	myGraph.PushFront(dC)
	myGraph.PushFront(dD)
	myGraph.PushFront(dE)
	myGraph.PushFront(dF)
	myGraph.PushFront(dG)
	myGraph.PushFront(dH)

	// A
	dA.AddNeighbor(dB, 5)
	dA.AddNeighbor(dH, 8)
	dA.AddNeighbor(dE, 9)

	// B
	dB.AddNeighbor(dC, 12)
	dB.AddNeighbor(dH, 4)
	dB.AddNeighbor(dD, 15)

	// C
	dC.AddNeighbor(dD, 3)
	dC.AddNeighbor(dG, 11)

	// D
	dD.AddNeighbor(dG, 9)

	// E
	dE.AddNeighbor(dF, 4)
	dE.AddNeighbor(dG, 20)
	dE.AddNeighbor(dH, 5)

	// F
	dF.AddNeighbor(dC, 1)
	dF.AddNeighbor(dG, 13)

	// G
	// Nada

	// H
	dH.AddNeighbor(dF, 6)
	dH.AddNeighbor(dC, 7)

	distances := dijkstra.ShortestPath(dA)
	vals := distances.Vals
	for a, d := range vals {
		println(a.Data.(string), d)
	}

}
