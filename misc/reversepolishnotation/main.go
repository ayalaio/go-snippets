package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ADD = "+"
	MIN = "-"
	DIV = "/"
	MUL = "*"
)

func isOp(t string) bool {

	for _, op := range []string{ADD, MIN, DIV, MUL} {
		if op == t {
			return true
		}
	}

	return false

}

func pop(stackPtr *[]int) int {
	var e int
	stack := *stackPtr
	e, *stackPtr = stack[len(stack)-1], stack[0:len(stack)-1]
	return e
}

func push(stackPtr *[]int, e int) {
	stack := *stackPtr
	*stackPtr = append(stack, e)
}

func reversePolish(a string) (int, error) {

	tokens := strings.Split(a, " ")

	stack := []int{}

	for _, token := range tokens {
		if isOp(token) {
			right := pop(&stack)
			left := pop(&stack)

			switch token {
			case ADD:
				push(&stack, left+right)
			case MIN:
				push(&stack, left-right)
			case MUL:
				push(&stack, left*right)
			case DIV:
				push(&stack, left/right)
			}
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("Not a valid string")
			}
			stack = append(stack, num)
		}
	}

	return stack[0], nil

}

func main() {
	var a string
	a = "15 7 1 1 + - / 3 * 2 1 1 + + -"

	num, err := reversePolish(a)

	if err != nil {
		fmt.Println("Not a valid string")
	} else {
		fmt.Println(num)
	}

}
