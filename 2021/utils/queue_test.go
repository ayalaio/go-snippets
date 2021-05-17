package utils

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	for e, i := q.Take(), 1; e != nil; e, i = q.Take(), i+1 {
		if e != i {
			t.Errorf("Expecting %d got %d", i, e)
		}
	}
}
