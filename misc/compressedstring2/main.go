package main

import (
	"fmt"
	"strconv"
)

func isNumber(str rune) bool {
	_, err := strconv.Atoi(string([]rune{str}))
	return err == nil
}

func expand(str string) string {

	for true {

		// greedly find first leaf parenthesis
		left := 0
		right := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '[' {
				left = i
			}
			if str[i] == ']' {
				right = i
				break // first leaf parenthesis found
			}
		}

		// exit if no other leaf parenthesis exist exit
		if left == 0 && right == 0 {
			break
		}

		// Find where the number of current parenthesis starts
		numPos := left - 1
		for numPos >= 0 && isNumber(rune(str[numPos])) {
			numPos--
		}
		numPos++

		// Convert that number to an integer
		quantToExpand, _ := strconv.Atoi(str[numPos:left])

		// expand the substring to that number
		expanded := ""
		for i := 0; i < quantToExpand; i++ {
			expanded += str[left+1 : right]
		}

		// replace all subset with the expansion
		str = string(str[:numPos]) + expanded + string(str[right+1:])

	}

	return str

}

func main() {
	a := "3[1[abc7[j]]4[xd]c]"

	b := expand(a)

	fmt.Println(b)

}
