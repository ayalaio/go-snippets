package main

type Node struct {
	Number int
	Left   *Node
	Right  *Node
}

func (n *Node) prune(currentSum, targetSum int) bool {

	currentSum += n.Number

	if n.Left == nil && n.Right == nil {
		if currentSum < targetSum {
			return true
		} else {
			return false
		}
	}

	pruneLeft, pruneRight := false, false

	if n.Left != nil {
		pruneLeft = n.Left.prune(currentSum, targetSum)
		if pruneLeft {
			println("prune", n.Left.Number)
			n.Left = nil
		}
	} else {
		pruneLeft = true
	}

	if n.Right != nil {
		pruneRight = n.Right.prune(currentSum, targetSum)
		if pruneRight {
			println("prune", n.Right.Number)
			n.Right = nil
		}
	} else {
		pruneRight = true
	}

	return pruneRight && pruneLeft
}

func main() {

	root := &Node{
		6,
		&Node{
			3, nil, nil,
		},
		&Node{
			8,
			&Node{
				4,
				&Node{
					1, nil, nil,
				},
				&Node{
					7, nil, nil,
				},
			},
			&Node{
				2,
				nil,
				&Node{
					3, nil, nil,
				},
			},
		},
	}

	root.prune(0, 20)

}
