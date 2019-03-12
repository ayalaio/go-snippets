package singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartLInstance(t *testing.T) {
	singleton := GetLInstance()
	sameSingleton := GetLInstance()

	n := 500

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go sameSingleton.AddOne()
	}

	fmt.Printf("Current coint is %d", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}

}
