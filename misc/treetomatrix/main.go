package main

import (
	"fmt"
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Loc struct {
	lvl, val int
}

type LocByOrder []Loc

func (self LocByOrder) Len() int {
	return len(self)
}
func (self LocByOrder) Less(i, j int) bool {
	if self[i].lvl == self[j].lvl {
		return self[i].val < self[j].val
	}
	return self[i].lvl < self[j].lvl
}
func (self LocByOrder) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func verticalTraversal(root *TreeNode) [][]int {

	paddingLeft := int(mostToTheLeft(root)) - 1   //remove root from count
	paddingRight := int(mostToTheRight(root)) - 1 //remove root from count

	// Rows to the left + root
	startingRow := paddingLeft + 1

	res := make([][]Loc, 0)

	// from first row to last row
	for i := 1; i <= paddingLeft+1+paddingRight; i++ {
		res = append(res, []Loc{})
	}

	// just remember rows start at idx 1
	verticalTraversalRec(root, startingRow, 0, res)

	r := [][]int{}
	for i := 0; i < len(res); i++ {
		row := res[i]
		sort.Sort(LocByOrder(row))
		r = append(r, []int{})
		for j := 0; j < len(res[i]); j++ {
			r[i] = append(r[i], row[j].val)
		}
	}

	return r

}

func mostToTheLeft(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(mostToTheLeft(node.Left)+1, mostToTheLeft(node.Right)-1)
}

func mostToTheRight(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(mostToTheRight(node.Left)-1, mostToTheRight(node.Right)+1)
}

func verticalTraversalRec(node *TreeNode, row, col int, res [][]Loc) {

	if node == nil {
		return
	}

	res[row-1] = append(res[row-1], Loc{col, node.Val})
	verticalTraversalRec(node.Left, row-1, col+1, res)
	verticalTraversalRec(node.Right, row+1, col+1, res)
}

func main() {

	var root *TreeNode

	root = &TreeNode{
		0,
		&TreeNode{
			1,
			&TreeNode{2, nil, nil},
			&TreeNode{3, nil, nil},
		},
		nil,
	}

	fmt.Println(verticalTraversal(root))

	root = &TreeNode{
		0,
		&TreeNode{
			1,
			nil,
			&TreeNode{2,
				&TreeNode{6, nil, nil},
				&TreeNode{
					3,
					nil,
					&TreeNode{
						4,
						nil,
						&TreeNode{5, nil, nil},
					},
				},
			},
		},
		nil,
	}

	fmt.Println(verticalTraversal(root))

}
