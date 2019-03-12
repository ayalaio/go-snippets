package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	goRoutines := 5
	wait.Add(goRoutines)
	for i := 0; i < goRoutines; i++ {
		go func(id int) {
			fmt.Printf("Hello from %d\n", id)
			wait.Done() // Same as wait.Add(-1)
		}(i)
	}
	wait.Wait()
}
