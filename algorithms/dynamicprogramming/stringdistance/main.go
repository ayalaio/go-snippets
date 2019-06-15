package main

import (
	"fmt"
	"math"
)

func minNum(i, j int, m [][]int) int {
	min := math.MaxInt64
	comp := []int{m[i-1][j-1], m[i][j-1], m[i-1][j]}
	for i := 0; i < len(comp); i++ {
		if comp[i] < min {
			min = comp[i]
		}
	}
	return min
}

func stringDistance(str1, str2 []rune) int {

	m := make([][]int, len(str1)+1) // +1 to consider empty string

	for i := 0; i < len(m); i++ {
		m[i] = make([]int, len(str2)+1) // +1 to consider empty string
		m[i][0] = i
	}

	for j := 0; j < len(m[0]); j++ {
		m[0][j] = j
	}

	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[0]); j++ {

			charIsChanged := str1[i-1] != str2[j-1]
			prevLowestChange := minNum(i, j, m)

			if charIsChanged {
				m[i][j] = prevLowestChange + 1
			} else {
				m[i][j] = prevLowestChange
			}
		}
	}
	fmt.Println(m)
	return m[len(m)-1][len(m[0])-1]

}

func main() {
	str1 := "azced"
	str2 := "abcdef"

	dst := stringDistance([]rune(str1), []rune(str2))

	fmt.Println(dst)
}
