package main

import "github.com/ayalaio/utils/stack"

type MathFact struct {
	Content int
}

func (self *MathFact) FactorialFunc() func(int) int {
	return func(j int) int {
		return self.Content * j
	}
}

func main() {
	s := stack.New()

	factorialOf := 6

	curr := factorialOf

	for curr >= 2 {
		s.Add(&MathFact{curr})
		curr--
	}

	// 1 -> base case
	r := 1
	for s.Len() > 0 {
		e := s.Front()
		ff := e.Value.(*MathFact).FactorialFunc()
		r = ff(r)
		s.Remove()
	}
	println(r)
}
