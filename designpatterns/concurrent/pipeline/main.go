package main

func generate(max int) (out chan int) {
	out = make(chan int)
	go (func() {
		for i := 1; i <= max; i++ {
			out <- i
		}
		close(out)
	})()
	return
}

func power2(in chan int) (out chan int) {
	out = make(chan int)

	go (func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	})()

	return
}

func sum(in chan int) (out chan int) {
	out = make(chan int)
	go (func() {
		var total int
		for i := range in {
			total += i
		}
		out <- total
		close(out)
	})()
	return
}

func main() {
	ch := sum(power2(generate(4)))
	println(<-ch)
}
