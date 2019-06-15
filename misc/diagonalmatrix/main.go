package main

import (
	"fmt"
	"math"
)

func translateDiagonal(n, lenY, lenX int) (int, int) {

	shifted := false

	if n > lenX*lenY/2 {
		n = lenX*lenY - n - 1
		shifted = true
	}

	k := int(math.Floor((-0.5 + math.Sqrt(0.25+2.0*float64(n)))))
	j := n - (k * (k + 1) / 2)
	i := k - j

	if shifted {
		i = lenY - i - 1
		j = lenX - j - 1
	}

	return i, j
}

func translateNormal(n, lenY, lenX int) (int, int) {
	i := n % lenX
	j := n / lenY
	return i, j
}

func traverseDiagonal2(m [][]int) {
	for k := 0; k < len(m)*len(m[0]); k++ {
		//translateDiagonal(k, len(m), len(m[0]))
		i, j := translateDiagonal(k, len(m), len(m[0]))
		fmt.Println(m[i][j])
	}
}

func traverseDiagonal(m [][]int) {

	for i := 0; i < len(m); i++ {
		for j := 0; j <= i; j++ {
			fmt.Println(m[i-j][j])
		}
	}

	for i := 1; i < len(m); i++ {
		for j := i; j < len(m); j++ {
			fmt.Println(m[i-j+len(m)-1][j])
		}
	}

}

func findNumberRec(m [][]int, target, sR, eR, sC, eC int) bool {

	if sR > eR {
		return false
	}

	if sC > eC {
		return false
	}

	midC := (sC + eC) / 2
	midR := (sR + eR) / 2

	if m[midR][midC] == target {
		return true
	} else if sR == midR && sC == midC {
		return false
	}

	if m[midR][midC] > target {

		return (findNumberRec(m, target, sR, midR, sC, midC) || //quadrant I
			findNumberRec(m, target, sR, midR, midC, eC) || //quadrant II
			findNumberRec(m, target, midR, eR, sC, midC)) //quadrant III

	} else {

		return (findNumberRec(m, target, sR, midR, midC, eC) || //quadrant II
			findNumberRec(m, target, midR, eR, sC, midC) || //quadrant III
			findNumberRec(m, target, midR, eR, midC, eC)) //quadrant IV

	}

}

func findNumber(m [][]int, target int) bool {
	return findNumberRec(m, target, 0, len(m), 0, len(m[0]))
}

func main() {
	m := [][]int{
		{1, 3, 6, 10, 15},
		{2, 5, 9, 14, 20},
		{4, 8, 13, 19, 24},
		{7, 12, 18, 23, 27},
		{11, 17, 22, 26, 29},
		{16, 21, 25, 28, 30},
	}

	_ = m

	//traverseDiagonal(m)
	traverseDiagonal2(m)

	fmt.Println(findNumber(m, 18))
	fmt.Println(findNumber(m, 17))
	fmt.Println(findNumber(m, 22))
	fmt.Println(findNumber(m, 35))

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(i, j)
	// }
}
