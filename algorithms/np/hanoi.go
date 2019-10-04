package main

import "fmt"

func hanoi(n, s, a, f rune) {

	if n == 1 {
		fmt.Printf("move from %c to %c\n", s, f)
		return
	}

	hanoi(n-1, s, f, a)
	fmt.Printf("move from %c to %c\n", s, f)
	hanoi(n-1, a, s, f)

}

func main() {
	hanoi(3, 'A', 'B', 'C')
}
