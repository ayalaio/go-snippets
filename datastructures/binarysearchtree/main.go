package main

import "fmt"

type BST struct {
	val         int
	Left, Right *BST
}

func NewBST(val int) *BST {
	return &BST{val, nil, nil}
}

func (self *BST) Add(val int) {
	if self.val > val {
		if self.Left != nil {
			self.Left.Add(val)
		} else {
			self.Left = NewBST(val)
			return
		}
	}

	if self.val <= val {
		if self.Right != nil {
			self.Right.Add(val)
		} else {
			self.Right = NewBST(val)
			return
		}
	}

}

func (self *BST) prevNode() *BST {
	if self.Left != nil {
		curr := self.Left
		for curr.Right != nil {
			curr = curr.Right
		}
		return curr
	}
	return nil
}

func (self *BST) Delete(val int, ptrToSelf **BST) {

	if self.val == val {

		if self.Left == nil && self.Right == nil {
			*ptrToSelf = nil
			return
		}

		if self.Left == nil {
			*ptrToSelf = self.Right
			return
		}

		if self.Right == nil {
			*ptrToSelf = self.Left
			return
		}

		prevNode := self.prevNode()
		self.val = prevNode.val

		self.Left.Delete(prevNode.val, &self.Left)

		return
	}

	if val >= self.val && self.Right != nil {
		self.Right.Delete(val, &self.Right)
		return
	}

	if val < self.val && self.Left != nil {
		self.Left.Delete(val, &self.Left)
		return
	}

}

func (self *BST) bfs() {

	clq := []*BST{self}
	nq := []*BST{}
	var e *BST

	for len(clq) > 0 {
		for len(clq) > 0 {
			e, clq = clq[0], clq[1:len(clq)]

			fmt.Print(e.val, " - ")

			if e.Left != nil {
				nq = append(nq, e.Left)
			}

			if e.Right != nil {
				nq = append(nq, e.Right)
			}

		}
		fmt.Println()
		clq = nq
		nq = []*BST{}
	}

}

func (self *BST) InOrder() {
	if self.Left != nil {
		self.Left.InOrder()
	}
	fmt.Println(self.val)
	if self.Right != nil {
		self.Right.InOrder()
	}
}

func main() {
	root := NewBST(5)
	root.Add(3)
	root.Add(7)
	root.Add(1)
	root.Add(4)
	root.Add(9)

	root.bfs()

	fmt.Println()

	root.Delete(5, &root)

	fmt.Println()
	root.bfs()
	root.InOrder()

}
