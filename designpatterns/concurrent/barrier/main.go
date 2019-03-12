package main

func Add(ch chan int, list []int) {
	l := len(list)
	if l == 1 {
		ch <- list[0]
	} else if l == 2 {
		ch <- list[0] + list[1]
	} else {
		var respChan chan int = make(chan int)
		go Add(respChan, list[:l/2])
		go Add(respChan, list[l/2:])
		a := <-respChan
		b := <-respChan
		ch <- a + b
		close(respChan)
	}
}

func main() {
	var respChan chan int = make(chan int)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	go Add(respChan, nums)
	println(<-respChan)
	close(respChan)
}
