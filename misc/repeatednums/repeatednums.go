package main

import (
	"fmt"
	"math"
)

func areThereRepeatedNums(b []int) bool {
	for i := 0; i < len(b); i++ {
		pos := int(math.Abs(float64(b[i])))
		if b[pos] < 0 {
			return true
		}
		b[pos] = -b[pos]
	}
	return false
}

func main() {
	a := []int{1, 7, 4, 7, 8, 9, 3, 6, 4, 7}

	fmt.Println(a)
	fmt.Println(areThereRepeatedNums(a))
	fmt.Println(a)
}
