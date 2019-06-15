package main

import (
	"fmt"
	"math"
)

type Queen struct {
	xPos, yPos int
	daughters  []*Queen
	mother     *Queen
}

func NewQueen(x, y int) *Queen {
	return &Queen{x, y, make([]*Queen, 0), nil}
}

type Board struct {
	root   *Queen
	size   int
	placed int
}

func (self *Queen) Colides(other *Queen) bool {
	if self.xPos == other.xPos {
		return true
	}

	if self.yPos == other.yPos {
		return true
	}

	if math.Abs(float64(self.yPos-other.yPos)) == math.Abs(float64(self.xPos-other.xPos)) {
		return true
	}

	return false

}

func (self *Queen) hasCollisions() bool {
	mom := self.mother
	for mom != nil {

		if self.Colides(mom) {
			return true
		}
		mom = mom.mother
	}

	return false
}

func (self *Queen) Populate(size int) (bool, *Queen) {
	if self.xPos == size-1 { // We reached the last position, this one is the final queen
		return true, self
	}
	for j := 0; j < size; j++ {
		nq := NewQueen(self.xPos+1, j)
		nq.mother = self
		if !nq.hasCollisions() {
			self.daughters = append(self.daughters, nq)
			ended, final := nq.Populate(size)
			if ended {
				return true, final
			}
		}
	}
	return false, nil
}

func (self *Queen) Print(mtx []string, lvl, size int) []string {

	for j := 0; j < size; j++ {
		if self.yPos == j {
			mtx[lvl] = mtx[lvl] + "Q"
		} else {
			mtx[lvl] = mtx[lvl] + "."
		}
	}
	mom := self.mother
	if mom != nil {
		return mom.Print(mtx, lvl+1, size)
	}
	return mtx
}

func main() {

	boardSize := 5

	solutions := [][]string{}

	for j := 0; j < boardSize; j++ {
		board := make([]string, boardSize)
		root := NewQueen(0, j)
		ended, final := root.Populate(boardSize)
		if ended {
			final.Print(board, 0, boardSize)
			solutions = append(solutions, board)
		}
	}

	fmt.Println(solutions)
}
