package main

import "fmt"

func findBitonic(a []int) []int {

	prevLong := make([]int, 0)
	currLong := make([]int, 0)

	wasIncreasing := a[0] < a[1]
	hasChangedBefore := false
	currLong = append(currLong, a[0])

	for i := 1; i < len(a); i++ {
		isIncreasing := a[i-1] < a[i]

		if wasIncreasing != isIncreasing {
			if hasChangedBefore {
				hasChangedBefore = false
				if len(currLong) > len(prevLong) {
					prevLong = currLong
				}
				currLong = make([]int, 0)
				i--
			} else {
				hasChangedBefore = true
			}
		}

		wasIncreasing = isIncreasing
		currLong = append(currLong, a[i])
	}

	if len(currLong) > len(prevLong) {
		prevLong = currLong
	}

	return prevLong
}

func main() {
	a := []int{1, 2, 3, 4, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1, 9, 0, 6, 8, 9, 4, 3, 0, 1, 2, 3, 7, 8, 9, 8, 6, 5, 4, 2, 1}
	b := findBitonic(a)
	fmt.Println(b)
}
