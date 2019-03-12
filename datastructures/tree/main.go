package main

import "fmt"

type Tree struct {
	node *Node
}

func (t *Tree)  insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
	return t
}


type Node struct {
	value int
	left *Node
	right *Node
}

func (n *Node) breathFirst() []int {
	if n == nil {
		return []int{}
	}
	nums := []int{n.value}
	nums = append(nums, n.left.breathFirst()...)
	nums = append(nums, n.right.breathFirst()...)
	return nums
}

func (n *Node) preOrder() []int {
	if n == nil {
		return []int{}
	}
	nums := n.left.preOrder()
	nums = append(nums, n.value)
	nums = append(nums, n.right.preOrder()...)
	return nums
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func (n *Node) insertNode(newNode *Node) {
	if newNode.value <= n.value {
		if n.left == nil {
			n.left = newNode
		} else {
			n.left.insertNode(newNode)
		}
	} else {
		if n.right == nil {
			n.right = newNode
		} else {
			n.right.insertNode(newNode)
		}
	}
}

func (n *Node) del(value int) {
	if value < n.value && n.left != nil {
		if n.left.value == value {
			rightRem := n.left.right
			n.left = n.left.left
			n.left.insertNode(rightRem)
		} else {
			 n.left.del(value)
		}
	} else if value > n.value && n.right != nil {
		if n.right.value == value {
			leftRem := n.right.left
			n.right = n.right.right
			n.right.insertNode(leftRem)
		} else {
			n.right.del(value)
		}
	}
}

func (n *Node) WhichFloor(value int) int {
	if n == nil {
		return -999999
	}
	floor := 0
	if value > n.value {
		floor = 1 + n.right.WhichFloor(value)
	} else if value < n.value {
		floor = 1 + n.left.WhichFloor(value)
	} // else floor is 0

	return floor
}


func main() {
	t := Tree{}
	t.insert(10).insert(8).insert(9).insert(4).insert(34)
	t.node.del(8)
	fmt.Println(t.node.preOrder())
	fmt.Println(t.node.breathFirst())
	fmt.Println(t.node.WhichFloor(34))
}
