package main

import (
	"fmt"
	"errors"
)

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node{
	return l.head
}

func (l *List) Last() *Node{
	return l.tail
}

func (l *List) Push(value int){
	node := &Node{value: value}
	if l.head == nil{
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}

	l.tail = node
}

func (l *List) Find(value int) (*Node, error) {
	n := l.First()
	err := errors.New(
		fmt.Sprintf("No such value %d present", value),
	)
	for {
		if n == nil {
			break
		}
		if n.value == value {
			err = nil
			break
		}
		n = n.Next()
		
	} 
	return n, err
}

func (l *List) StringForward() string {
	str := fmt.Sprint("[")
	n := l.First()
	for {
		if n == nil {
			break
		}
		str = fmt.Sprint(str, n.value, " ")
		n = n.Next()
	}
	return fmt.Sprint(str, "]")
}

func (l *List) StringBackward() string {
	str := fmt.Sprint("[")
	n := l.Last()
	for {
		if n == nil {
			break
		}
		str = fmt.Sprint(str, n.value, " ")
		n = n.Prev()
	}
	return fmt.Sprint(str, "]")
}

func (l *List) String() string {
	return l.StringForward()
}


type Node struct {
	value int
	next *Node
	prev *Node
}

func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) String() string {
	return fmt.Sprint("<", n.value, ">")
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	fmt.Println(l.StringForward())
	fmt.Println(l.StringBackward())
	if n, err := l.Find(7); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Node ", n , " was found")
	}
}
