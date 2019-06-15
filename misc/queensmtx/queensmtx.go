package main

import "fmt"

func printBoard(board [][]rune) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func safe(board [][]rune, row, col int) bool {

	// Vertical
	for i := 0; i < len(board); i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Horizontal
	for j := 0; j < len(board); j++ {
		if board[row][j] == 'Q' {
			return false
		}
	}

	// Diagonal \
	// for i, j := row+1, col+1; i < len(board) && j < len(board); i, j = i+1, j+1 {
	// 	if board[i][j] == 'Q' {
	// 		return false
	// 	}
	// }
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Diagonal \
	// for i, j := row+1, col-1; i < len(board) && j >= 0; i, j = i+1, j-1 {
	// 	if board[i][j] == 'Q' {
	// 		return false
	// 	}
	// }
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func nQueens(board [][]rune, lvl int) {

	if len(board) == lvl {
		printBoard(board)
		return
	}

	for j := 0; j < len(board[lvl]); j++ {
		if safe(board, lvl, j) {
			board[lvl][j] = 'Q'
			nQueens(board, lvl+1)
			board[lvl][j] = '-'
		}
	}
}

func main() {

	board := [][]rune{
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
	}

	nQueens(board, 0)

}
