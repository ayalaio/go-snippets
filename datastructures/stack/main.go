package main

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}
func (s *Stack) Pop() int {
	lastPosition := len(s.items) - 1
	if lastPosition == -1{
		return -1
	}
	item := s.items[lastPosition]
	s.items = s.items[0:lastPosition]
	return item
}

func main() {
	s := Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
}
