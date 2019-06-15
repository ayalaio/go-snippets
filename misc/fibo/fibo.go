package main

func main() {

	fiboOf := 5

	prevA := 1
	prevB := 1

	var res int

	for i := 1; i <= fiboOf; i++ {
		res = prevA + prevB
		prevA = prevB
		prevB = res
	}

	println(res)

}
