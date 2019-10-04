package main

import (
	"fmt"
	"math"
)

func distance(str1, str2 string) int {
	return distanceR(str1, str2, len(str1)-1, len(str2)-1)
}

func distanceR(str1, str2 string, iE, jE int) int {
	if iE == 0 {
		return jE
	}

	if jE == 0 {
		return iE
	}

	// Adding a character
	aac := 1 + distanceR(str1, str2, iE, jE-1)

	// Removing a character
	rac := 1 + distanceR(str1, str2, iE-1, jE)

	// Changing a character
	dist := 1
	if str1[iE-1] == str2[jE-1] {
		dist = 0
	}
	cac := dist + distanceR(str1, str2, iE-1, jE-1)

	min := math.MaxInt64
	for _, dist := range []int{aac, rac, cac} {
		if dist < min {
			min = dist
		}
	}
	return min
}

func main() {
	d := distance("olaarason", "holacorazon")
	fmt.Println(d)
}
