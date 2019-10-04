package main

import (
	"fmt"
	"math"
)

type MinHeap struct {
	list []int
	size int
}

func NewMinHeap() *MinHeap {
	a := make([]int, 100)
	for i := 0; i < len(a); i++ {
		a[i] = math.MaxInt64
	}
	return &MinHeap{a, 0}
}

func (self *MinHeap) peek() int {
	return self.list[0]
}

func (self *MinHeap) insert(v int) {

	if self.size == 0 {
		self.list[0] = v
		self.size = 1
	} else {
		self.list[self.size] = v
		self.size++
		self.bubbleUp(self.size - 1)
	}

}

func (self *MinHeap) delete() {
	if self.size == 1 {
		self.list[0] = math.MaxInt64
		self.size--
	} else {
		self.list[0] = self.list[self.size-1]
		self.list[self.size-1] = math.MaxInt64
		self.size--
		self.bubbleDown(0)
	}
}

// if lower, go upper
func (self *MinHeap) bubbleUp(pos int) {

	if pos == 0 {
		return
	}

	current := self.list[pos]
	parent := self.list[Parent(pos)]

	if current < parent {
		self.list[pos] = parent
		self.list[Parent(pos)] = current
		self.bubbleUp(Parent(pos))
	}

}

// if bigger, go downer
func (self *MinHeap) bubbleDown(pos int) {

	if pos >= self.size {
		return
	}

	current := self.list[pos]
	childR := self.list[RightChild(pos)]
	childL := self.list[LeftChild(pos)]

	var child, childPos int
	if childR < childL {
		child = childR
		childPos = RightChild(pos)
	} else {
		child = childL
		childPos = LeftChild(pos)
	}

	if current > child {
		self.list[childPos] = current
		self.list[pos] = child
		self.bubbleDown(childPos)
	}

}

func LeftChild(pos int) int {
	return pos*2 + 1
}

func RightChild(pos int) int {
	return pos*2 + 2
}

func Parent(pos int) int {
	if pos == 0 {
		return 0
	}
	return (pos - 1) / 2
}

func main() {
	h := NewMinHeap()
	h.insert(4)
	fmt.Println(h.peek())
	h.insert(7)
	fmt.Println(h.peek())
	h.insert(2)
	fmt.Println(h.peek())
	h.insert(9)
	fmt.Println(h.peek())
	h.insert(10)
	fmt.Println(h.peek())

	fmt.Println("deleting")
	fmt.Println(h.peek())
	h.delete()

	fmt.Println(h.peek())
	h.delete()

	fmt.Println(h.peek())
	h.delete()

	fmt.Println(h.peek())
	h.delete()

	fmt.Println(h.peek())
	h.delete()

}
