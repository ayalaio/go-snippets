package utils

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	for e, i := s.Pop(), 3; e != nil; e, i = s.Pop(), i-1 {
		if e != i {
			t.Errorf("Expecting %d got %d", i, e)
		}
	}
}
