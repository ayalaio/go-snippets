package cycledetection

import (
	"fmt"

	"github.com/ayalaio/utils/vertex"
)

// FindCycle finds a cycle in a Graph
func FindCycle(v *vertex.Vertex) bool {
	isBeingVisited := make(map[*vertex.Vertex]bool)

	var dfs func(v *vertex.Vertex) bool
	dfs = func(v *vertex.Vertex) bool {
		fmt.Printf("%d\n", v.Data)

		isBeingVisited[v] = true
		v.Visit()

		neighborHasCycles := false

		for vne := v.Neighbors.Front(); vne != nil; vne = vne.Next() {
			vn := vne.Value.(*vertex.Vertex)
			if isBeingVisited[vn] {
				return true
			}

			if !vn.IsVisited && !neighborHasCycles {
				neighborHasCycles = dfs(vn)
			}
		}

		isBeingVisited[v] = false
		return neighborHasCycles
	}

	return dfs(v)
}
