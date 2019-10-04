package main

import "fmt"

type List struct {
	Start, End *Node
}

func NewList() *List {
	return &List{}
}

func (self *List) Add(val int) {
	n := NewNode(val)
	if self.Start == nil {
		self.Start = n
		self.End = n
	} else {
		self.End.Next = n
		self.End = n
	}
}

func (self *List) Delete(val int) {
	var prevP *Node = nil
	p := self.Start
	for p != nil {
		if p.Val == val {
			if prevP == nil {
				self.Start = p.Next
			} else {
				prevP.Next = p.Next
			}

			if self.End == p {
				if prevP.Next != nil {
					self.End = prevP.Next
				} else {
					self.End = prevP
				}

			}
		}
		prevP = p
		p = p.Next
	}
}

func (self *List) Reverse() {
	var prevP, nextP *Node
	p := self.Start

	for p != nil {
		nextP = p.Next
		p.Next = prevP

		prevP = p
		p = nextP
	}

	self.Start, self.End = self.End, self.Start

}

func (self *List) String() string {
	p := self.Start
	s := ""
	for p != nil {
		s = fmt.Sprintf("%s-%d", s, p.Val)
		p = p.Next
	}
	return s
}

type Node struct {
	Val  int
	Next *Node
}

func NewNode(v int) *Node {
	return &Node{v, nil}
}

func main() {
	a := NewList()
	a.Add(1)
	a.Add(2)
	a.Add(4)
	a.Add(3)
	fmt.Println(a)
	a.Delete(4)
	fmt.Println(a)
	a.Add(4)
	a.Add(5)
	a.Reverse()
	fmt.Println(a)

}
