package dijkstra

import (
	"container/heap"
)

type Distances struct {
	Vals     map[*Vertex]int
	defaults int
}

func NewDistances() Distances {
	return Distances{make(map[*Vertex]int), 9999}
}

func (d *Distances) Get(v *Vertex) int {
	if val, ok := d.Vals[v]; ok {
		return val
	}
	return d.defaults
}

func (d *Distances) Set(v *Vertex, w int) {
	d.Vals[v] = w
}

type Precendents struct {
	vals map[*Vertex]*Vertex
}

func NewPrecendents() Precendents {
	return Precendents{make(map[*Vertex]*Vertex)}
}

func (d *Precendents) Get(v *Vertex) *Vertex {
	return d.vals[v]
}

func (d *Precendents) Set(u, v *Vertex) {
	d.vals[u] = v
}

type EdgeHeap []Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func ShortestPath(root *Vertex) Distances {

	distances := NewDistances()
	precedents := NewPrecendents()

	q := &EdgeHeap{}
	heap.Init(q)

	heap.Push(q, Edge{root, 0})
	distances.Set(root, 0)

	for q.Len() > 0 {

		// fmt.Printf("\n\n")
		// fmt.Println(q)

		edge := heap.Pop(q).(Edge)
		curr := edge.Target
		currWeight := distances.Get(curr)

		// fmt.Printf("Current %s: %d\n", curr.Data, currWeight)
		for edgeElement := curr.Neighbors.Front(); edgeElement != nil; edgeElement = edgeElement.Next() {
			edgeToN := edgeElement.Value.(Edge)
			weightToN := edgeToN.Weight
			N := edgeToN.Target

			// fmt.Printf("Neighbor %s: %d\n", N.Data, weightToN)

			if distances.Get(N) > currWeight+weightToN {
				distances.Set(N, currWeight+weightToN)
				precedents.Set(N, curr)
			}

			heap.Push(q, edgeToN)
		}
	}

	return distances

}
