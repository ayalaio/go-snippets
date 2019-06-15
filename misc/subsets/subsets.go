package main

import "fmt"

type SubSetHolder struct {
	left    *SubSetHolder
	right   *SubSetHolder
	content []int
}

func insertForSubSet(node *SubSetHolder, n int) {
	content := node.content
	if node.left == nil && node.right == nil {
		node.left = &SubSetHolder{nil, nil, append(content, n)}
		node.right = &SubSetHolder{nil, nil, content}
		return
	}
	insertForSubSet(node.left, n)
	insertForSubSet(node.right, n)
}

func buildSubSets(a []int) *SubSetHolder {

	root := &SubSetHolder{nil, nil, make([]int, 0)}

	for i := 0; i < len(a); i++ {
		insertForSubSet(root, a[i])
	}

	return root

}

func printTree(node *SubSetHolder) {
	if node.left == nil && node.right == nil {
		fmt.Println(node.content)
	} else {
		printTree(node.left)
		printTree(node.right)
	}

}

func printSubArrays(a []int) {
	for i := 0; i < len(a); i++ {
		for k := i; k < len(a); k++ {
			for j := i; j <= k; j++ {
				fmt.Print(a[j])
			}
			fmt.Println()
		}
	}
}

func backtrackSubSet(original, res []int, level int) {

	if level == len(res) {
		fmt.Println(res)
		return
	}

	elem := original[level]

	// add

	backtrackSubSet(original, append(res, elem), 1)

	// no-add

	backtrackSubSet(original, res, 1)
}

func main() {
	a := []int{1, 2, 3}

	printTree(buildSubSets(a))

	backtrackSubSet(a, make([]int, 0), 0)

	fmt.Println()

	printSubArrays(a)

}
