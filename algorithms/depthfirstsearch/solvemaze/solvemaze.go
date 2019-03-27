package solvemaze

import (
	"container/list"
)

func Find(mtrx [][]int) *list.List {

	var startPointX, startPointY int

	for i := 0; i < len(mtrx); i++ {
		for j := 0; j < len(mtrx[0]); j++ {
			if mtrx[i][j] == 2 {
				startPointX, startPointY = i, j
			}
		}
	}

	var endPointX, endPointY int

	for i := 0; i < len(mtrx); i++ {
		for j := 0; j < len(mtrx[0]); j++ {
			if mtrx[i][j] == 3 {
				endPointX, endPointY = i, j
			}
		}
	}

	m := &maze{
		start:     &Coordinate{startPointX, startPointY},
		end:       &Coordinate{endPointX, endPointY},
		mtrx:      mtrx,
		visitMtrx: newVisitMatrix(len(mtrx[0]), len(mtrx)),
	}

	return m.find()
}

func newVisitMatrix(sizeX, sizeY int) [][]bool {
	a := make([][]bool, sizeY)
	for i := range a {
		a[i] = make([]bool, sizeX)
	}
	return a
}

type Coordinate struct {
	X, Y int
}

type maze struct {
	start     *Coordinate
	end       *Coordinate
	mtrx      [][]int
	visitMtrx [][]bool
}

func (m *maze) find() *list.List {
	l, found := m.dfs(m.start.X, m.start.Y, *list.New())
	if found {
		return l
	}
	return nil
}

func (m *maze) dfs(i, j int, _l list.List) (*list.List, bool) {

	var l *list.List
	l = &_l

	if i < 0 || j < 0 || i >= len(m.mtrx) || j >= len(m.mtrx[0]) { // Out of bounds
		return l, false
	}

	if m.visitMtrx[i][j] { // Already visited
		return l, false
	}

	if m.mtrx[i][j] == 1 { // Wall
		return l, false
	}

	// In reality this "list" is a "tree".
	// Where branches go to the same root
	l.PushFront(&Coordinate{i, j})

	m.visitMtrx[i][j] = true

	if i == m.end.X && j == m.end.Y {
		return l, true
	}

	l1, f1 := m.dfs(i+1, j, *l)
	if f1 {
		return l1, true
	}

	l2, f2 := m.dfs(i, j+1, *l)
	if f2 {
		return l2, true
	}

	l3, f3 := m.dfs(i-1, j, *l)
	if f3 {
		return l3, true
	}

	l4, f4 := m.dfs(i, j-1, *l)
	if f4 {
		return l4, true
	}

	return nil, false
}
