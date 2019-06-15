package main

import "fmt"

type Hamilton struct {
	selectedEdges      map[*Edge]int
	orderSelectedEdges []*Edge
	visitedVertex      map[*Vertex]bool
}

type Vertex struct {
	val   string
	edges []*Edge
}

func NewVertex(val string) *Vertex {
	return &Vertex{val, make([]*Edge, 0)}
}

type Edge struct {
	v, u *Vertex
}

func (e *Edge) Name() string {
	return fmt.Sprint("(", e.u.val, ",", e.v.val, ")")
}

func NewEdge(v, u *Vertex) *Edge {
	return &Edge{v, u}
}

func (self *Vertex) AddNeighbor(v *Vertex) {
	e := NewEdge(self, v)
	self.edges = append(self.edges, e)
	v.edges = append(v.edges, e)
}

func (self *Hamilton) solve(node *Vertex, lvl int) bool {
	if _, ok := self.visitedVertex[node]; ok {
		return false
	}
	self.visitedVertex[node] = true

	if len(self.visitedVertex) == len(self.orderSelectedEdges)+1 {
		return true
	}

	for i := 0; i < len(node.edges); i++ {
		e := node.edges[i]
		if _, ok := self.selectedEdges[e]; ok {
			continue
		}
		self.selectedEdges[e] = lvl
		self.orderSelectedEdges[lvl] = e
		var solved bool
		if node == e.u {
			solved = self.solve(e.v, lvl+1)
		} else {
			solved = self.solve(e.u, lvl+1)
		}
		if solved {
			return true
		}
		delete(self.selectedEdges, e)
		self.orderSelectedEdges[lvl] = nil
	}
	delete(self.visitedVertex, node)

	return false
}

func NewHamilton(totalNumOfVertex int) *Hamilton {
	return &Hamilton{
		selectedEdges:      make(map[*Edge]int),
		orderSelectedEdges: make([]*Edge, totalNumOfVertex-1),
		visitedVertex:      make(map[*Vertex]bool),
	}
}

func (self *Hamilton) Print() {
	for idx, e := range self.orderSelectedEdges {
		fmt.Print(e.Name(), "@", idx, " :")
	}
}

func Solve(nodes []*Vertex) {

	totalNumOfVertex := len(nodes)

	for i := 0; i < len(nodes); i++ {
		hamilton := NewHamilton(totalNumOfVertex)
		hasSolution := hamilton.solve(nodes[i], 0)
		if hasSolution {
			hamilton.Print()
		}
		fmt.Println()
	}
}

func main() {
	var a, b, c, d, e, f, g, h, i, j, k *Vertex
	a = NewVertex("a")
	b = NewVertex("b")
	c = NewVertex("c")
	d = NewVertex("d")
	e = NewVertex("e")
	f = NewVertex("f")
	g = NewVertex("g")
	h = NewVertex("h")
	i = NewVertex("i")
	j = NewVertex("j")
	k = NewVertex("k")

	a.AddNeighbor(b)
	a.AddNeighbor(i)

	b.AddNeighbor(c)
	b.AddNeighbor(d)

	i.AddNeighbor(g)
	i.AddNeighbor(h)

	c.AddNeighbor(d)
	h.AddNeighbor(g)

	d.AddNeighbor(e)
	g.AddNeighbor(f)

	f.AddNeighbor(e)
	f.AddNeighbor(j)

	k.AddNeighbor(e)

	Solve([]*Vertex{a, b, c, d, e, f, g, h, i, j, k})
}
