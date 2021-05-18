package backtracking

import (
	"testing"
)

func hasCollision(board [][]bool, x, y int) bool {
	// up and down
	for i := 0; i < len(board); i++ {
		if board[i][y] && x != i {
			return true
		}
	}

	// left and right
	for j := 0; j < len(board); j++ {
		if board[x][j] && y != j {
			return true
		}
	}

	// inverted slash
	// down
	for i, j := x, y; i < len(board) && j < len(board[0]) && i >= 0 && j >= 0; i, j = i+1, j+1 {
		if board[i][j] && x != i && y != j {
			return true
		}
	}
	// up
	for i, j := x, y; i < len(board) && j < len(board[0]) && i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] && x != i && y != j {
			return true
		}
	}

	// forward slash
	// down
	for i, j := x, y; i < len(board) && j < len(board[0]) && i >= 0 && j >= 0; i, j = i+1, j-1 {
		if board[i][j] && x != i && y != j {
			return true
		}
	}
	// up
	for i, j := x, y; i < len(board) && j < len(board[0]) && i >= 0 && j >= 0; i, j = i-1, j+1 {
		if board[i][j] && x != i && y != j {
			return true
		}
	}

	return false
}

func tryAddQueen(board [][]bool, x, y int) bool {
	if !hasCollision(board, x, y) {
		board[x][y] = true
		return true
	}
	return false
}

func SolveQueensInBoard(level int, board [][]bool) bool {
	hasLowerLevelSolved := false
	for j := 0; j < len(board); j++ {
		if tryAddQueen(board, level, j) {
			if level == len(board)-1 {
				return true
			}
			if SolveQueensInBoard(level+1, board) {
				hasLowerLevelSolved = true
			} else {
				board[level][j] = false
			}
		}
	}
	return hasLowerLevelSolved
}

func Test_SolveQueensInBoard(t *testing.T) {
	size := 15
	board := make([][]bool, size)
	for i := range board {
		board[i] = make([]bool, size)
	}

	SolveQueensInBoard(0, board)
	numOfQueensInBoard := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				numOfQueensInBoard++
				if hasCollision(board, i, j) {
					t.Errorf("Collision detected at %d,%d", i, j)
				}
			}
		}
	}
	if numOfQueensInBoard != size {
		t.Errorf("We were expecting %d queens, just got %d", size, numOfQueensInBoard)
	}
}
