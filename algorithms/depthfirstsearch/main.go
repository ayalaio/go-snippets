package main

import (
	"fmt"

	"github.com/ayalaio/go-snippets/algorithms/depthfirstsearch/cycledetection"
	"github.com/ayalaio/go-snippets/algorithms/depthfirstsearch/dfs"
	"github.com/ayalaio/go-snippets/algorithms/depthfirstsearch/solvemaze"
	"github.com/ayalaio/go-snippets/algorithms/depthfirstsearch/topologicalorder"
	"github.com/ayalaio/utils/vertex"
)

func main() {

	println("------ dfs ------")
	l := []*vertex.Vertex{}

	l = append(l, vertex.NewVertex(0))
	l = append(l, vertex.NewVertex(1))
	l = append(l, vertex.NewVertex(2))
	l = append(l, vertex.NewVertex(3))
	l = append(l, vertex.NewVertex(4))

	l[1].AddNeighbor(l[0])
	l[0].AddNeighbor(l[3])
	l[1].AddNeighbor(l[2])
	l[3].AddNeighbor(l[4])

	dfs.DFS(l[1:2])

	println("\n------ dfs rec ------")
	l = []*vertex.Vertex{}

	l = append(l, vertex.NewVertex(0))
	l = append(l, vertex.NewVertex(1))
	l = append(l, vertex.NewVertex(2))
	l = append(l, vertex.NewVertex(3))
	l = append(l, vertex.NewVertex(4))

	l[1].AddNeighbor(l[0])
	l[0].AddNeighbor(l[3])
	l[1].AddNeighbor(l[2])
	l[3].AddNeighbor(l[4])

	dfs.DFSRec(l[1:2])

	println("\n------ topologicalorder ------")
	l = []*vertex.Vertex{}

	l = append(l, vertex.NewVertex(0))
	l = append(l, vertex.NewVertex(1))
	l = append(l, vertex.NewVertex(2))
	l = append(l, vertex.NewVertex(3))
	l = append(l, vertex.NewVertex(4))

	l[1].AddNeighbor(l[0])
	l[0].AddNeighbor(l[3])
	l[1].AddNeighbor(l[2])
	l[3].AddNeighbor(l[4])

	t := topologicalorder.New()
	order := t.GetOrder(l)

	for e := order.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d\n", e.Value.(*vertex.Vertex).Data)
	}

	println("\n------ cycle detection ------")

	v1 := vertex.NewVertex(1)
	v2 := vertex.NewVertex(2)
	v3 := vertex.NewVertex(3)
	v4 := vertex.NewVertex(4)
	v5 := vertex.NewVertex(5)
	v6 := vertex.NewVertex(6)

	v1.AddNeighbor(v2)
	v2.AddNeighbor(v3)
	v2.AddNeighbor(v4)
	v3.AddNeighbor(v5)
	v4.AddNeighbor(v5)
	v4.AddNeighbor(v6)
	//v5.AddNeighbor(v1)

	if !cycledetection.FindCycle(v1) {
		fmt.Printf("There is no cycle \n")
	}

	v1 = vertex.NewVertex(1)
	v2 = vertex.NewVertex(2)
	v3 = vertex.NewVertex(3)
	v4 = vertex.NewVertex(4)
	v5 = vertex.NewVertex(5)
	v6 = vertex.NewVertex(6)

	v1.AddNeighbor(v2)
	v2.AddNeighbor(v3)
	v2.AddNeighbor(v4)
	v3.AddNeighbor(v5)
	v4.AddNeighbor(v5)
	v4.AddNeighbor(v6)
	v6.AddNeighbor(v1)

	if cycledetection.FindCycle(v1) {
		fmt.Printf("There is a cycle \n")
	}

	println("------ solvemaze -------")

	a := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 2, 1},
		{1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1},
		{1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1},
		{1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 3, 0, 0, 0, 1, 1},
	}

	lst := solvemaze.Find(a)
	if l != nil {
		fmt.Printf("maze solution found %d\n", lst.Len())
		for e := lst.Front(); e != nil; e = e.Next() {
			if e.Value != nil {
				c := e.Value.(*solvemaze.Coordinate)
				println(c.X, c.Y)
			}
		}
	}

}
