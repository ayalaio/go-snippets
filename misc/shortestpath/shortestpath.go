package main

import (
	"fmt"
	"math"
)

func solveL(a, v [][]int, startR, startC, finalR, finalC, dist int) (int, error) {

	if startR == finalR && startC == finalC {
		return dist, nil
	}

	if startR < 0 {
		return 0, fmt.Errorf("Out of bounds")
	} else if startR >= len(a) {
		return 0, fmt.Errorf("Out of bounds")
	} else if startC < 0 {
		return 0, fmt.Errorf("Out of bounds")
	} else if startC >= len(a[startR]) {
		return 0, fmt.Errorf("Out of bounds")
	} else if a[startR][startC] == 0 {
		return 0, fmt.Errorf("A wall")
	} else if v[startR][startC] > 0 {
		return 0, fmt.Errorf("Area already visted")
	}

	v[startR][startC] = dist

	rA, _ := solveL(a, v, startR+1, startC, finalR, finalC, dist+1)
	rB, _ := solveL(a, v, startR-1, startC, finalR, finalC, dist+1)
	rC, _ := solveL(a, v, startR, startC+1, finalR, finalC, dist+1)
	rD, _ := solveL(a, v, startR, startC-1, finalR, finalC, dist+1)

	// backtracking
	v[startR][startC] = 0

	max := 0

	for _, n := range []int{rA, rB, rC, rD} {
		if n > max {
			max = n
		}
	}

	return max, nil

}

func solve(a, v [][]int, startR, startC, finalR, finalC, dist int) (int, error) {

	if startR == finalR && startC == finalC {
		return dist, nil
	}

	if startR < 0 {
		return math.MaxInt64, fmt.Errorf("Out of bounds")
	} else if startR >= len(a) {
		return math.MaxInt64, fmt.Errorf("Out of bounds")
	} else if startC < 0 {
		return math.MaxInt64, fmt.Errorf("Out of bounds")
	} else if startC >= len(a[startR]) {
		return math.MaxInt64, fmt.Errorf("Out of bounds")
	} else if a[startR][startC] == 0 {
		return math.MaxInt64, fmt.Errorf("A wall")
	} else if v[startR][startC] > 0 && v[startR][startC] < dist {
		return math.MaxInt64, fmt.Errorf("Area already visted")
	}

	v[startR][startC] = dist

	rA, _ := solve(a, v, startR+1, startC, finalR, finalC, dist+1)
	rB, _ := solve(a, v, startR-1, startC, finalR, finalC, dist+1)
	rC, _ := solve(a, v, startR, startC+1, finalR, finalC, dist+1)
	rD, _ := solve(a, v, startR, startC-1, finalR, finalC, dist+1)

	min := math.MaxInt64

	for _, n := range []int{rA, rB, rC, rD} {
		if n < min {
			min = n
		}
	}

	return min, nil

}

func findShortestPath(a [][]int) (int, error) {

	initC := 0
	initR := 0
	finlC := 0
	finlR := 0
	v := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		v[i] = make([]int, len(a[i]))
		for j := 0; j < len(a[0]); j++ {
			v[i][j] = 0
			if a[i][j] == 2 {
				initR, initC = i, j
			}
			if a[i][j] == 3 {
				finlR, finlC = i, j
			}
		}
	}

	return solve(a, v, initR, initC, finlR, finlC, 0)

}

func findLongestPath(a [][]int) (int, error) {

	initC := 0
	initR := 0
	finlC := 0
	finlR := 0
	v := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		v[i] = make([]int, len(a[i]))
		for j := 0; j < len(a[0]); j++ {
			v[i][j] = 0
			if a[i][j] == 2 {
				initR, initC = i, j
			}
			if a[i][j] == 3 {
				finlR, finlC = i, j
			}
		}
	}

	r, s := solveL(a, v, initR, initC, finlR, finlC, 0)

	// for i := 0; i < len(v); i++ {
	// 	for j := 0; j < len(v); j++ {
	// 		println(v[i][j])
	// 	}
	// }

	return r, s

}

func main() {

	a := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 3, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0},
	}

	b, err := findShortestPath(a)

	if err == nil {
		fmt.Println(b)
	}

	b, err = findLongestPath(a)

	if err == nil {
		fmt.Println(b)
	}

}
