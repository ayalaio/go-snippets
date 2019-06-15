package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type HolderNode struct {
	Val    *TreeNode
	Parent *HolderNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	a1 := getPath(root, p, nil)
	a2 := getPath(root, q, nil)

	as1 := holderNodeSize(a1)
	as2 := holderNodeSize(a2)

	asT := as1 - as2
	if asT > 0 {
		a1 = levelHolder(a1, asT)
	} else if asT < 0 {
		a2 = levelHolder(a2, -asT)
	}

	for a1.Val != a2.Val {
		a1 = levelHolder(a1, 1)
		a2 = levelHolder(a2, 1)
	}

	return a1.Val

}

func levelHolder(h *HolderNode, i int) *HolderNode {
	curr := h
	for i > 0 {
		curr = curr.Parent
		i--
	}
	return curr
}

func holderNodeSize(h *HolderNode) int {
	count := 0
	for i := h; i != nil; i = i.Parent {
		count++
	}
	return count
}

func getPath(src, target *TreeNode, hn *HolderNode) *HolderNode {

	hn = &HolderNode{src, hn}

	if src == target {
		return hn
	}

	if src.Left == nil && src.Right == nil {
		return nil
	}

	var path *HolderNode

	if path == nil && src.Left != nil {
		path = getPath(src.Left, target, hn)
	}

	if path == nil && src.Right != nil {
		path = getPath(src.Right, target, hn)
	}

	return path

}

func main() {
	p := &TreeNode{5, nil, nil}
	q := &TreeNode{1, nil, nil}

	root := &TreeNode{3,
		p,
		q,
	}

	r := lowestCommonAncestor(root, p, q)
	println(r)
}
