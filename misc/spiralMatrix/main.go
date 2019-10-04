package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{9, 10, 11, 12, 13, 14, 15, 16},
		{17, 18, 19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30, 31, 32},
		{33, 34, 35, 36, 37, 38, 39, 40},
	}

	printSpiral(a)
}

func printSpiral(a [][]int) {

	const (
		R = 1
		B = 2
		L = 3
		T = 4
	)

	limitL := 0
	limitR := len(a[0]) - 1
	limitT := 0
	limitB := len(a) - 1

	moving := R

	i, j := 0, 0

	for limitB >= limitT && limitR >= limitL {

		if moving == R && j == limitR {
			moving = B
			limitT++

		} else if moving == B && i == limitB {
			moving = L
			limitR--

		} else if moving == L && j == limitL {
			moving = T
			limitB--

		} else if moving == T && i == limitT {
			moving = R
			limitL++

		}

		fmt.Println(a[i][j])

		if moving == R {
			j++
		}
		if moving == B {
			i++
		}
		if moving == L {
			j--
		}
		if moving == T {
			i--
		}

	}

}
