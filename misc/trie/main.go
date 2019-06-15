package main

import "fmt"

type Trie struct {
	substring string
	position  int
	children  []*Trie
}

func contains(newString, oldString string) int {
	pos := 0
	for pos < len(oldString) && pos < len(newString) && newString[pos] == oldString[pos] {
		pos++
	}
	return pos - 1
}

func (t *Trie) insertChild(child *Trie) {
	t.children = append(t.children, child)
}

func (root *Trie) insertSubstring(substring string, substringPos int) {
	childrenContainsSubstring := false
	for i := 0; i < len(root.children); i++ {
		node := root.children[i]
		pos := contains(substring, node.substring)
		if pos >= 0 {

			currString := node.substring
			currPosition := node.position

			node.substring = node.substring[:pos+1]
			node.position = -1

			childrenContainsSubstring = true
			node.insertSubstring(currString[pos+1:], currPosition)
			node.insertSubstring(substring[pos+1:], substringPos)
		}
	}
	if !childrenContainsSubstring {
		t := &Trie{substring, substringPos, make([]*Trie, 0)}
		root.insertChild(t)
	}
}

func bf(t *Trie) {
	q := []*Trie{t}
	var el *Trie
	for len(q) > 0 {
		el, q = q[0], q[1:]
		fmt.Println("child of", el.substring, el.position)
		for i := 0; i < len(el.children); i++ {
			fmt.Print(el.children[i].substring, el.children[i].position, "     ")
			q = append(q, el.children[i])
		}
		fmt.Println()

	}

}

func substringFromBoth(t *Trie) (string, bool, bool) {
	if t == nil {
		return "", false, false
	}

	// its a leaf trie
	if t.position > -1 {
		lastChar := t.substring[len(t.substring)-1]
		if lastChar == '#' {
			return "", false, true
		} else if lastChar == '$' {
			return "", true, false
		}
	}

	currLongest := ""
	currHasDollar := false
	currHasHash := false

	fmt.Println(t.substring)

	for i := 0; i < len(t.children); i++ {
		longest, hasDollar, hasHash := substringFromBoth(t.children[i])
		if hasDollar && hasHash {
			if len(longest) > len(currLongest) {
				currLongest = longest
			}
		}
		currHasDollar = currHasDollar || hasDollar
		currHasHash = currHasHash || hasHash
	}

	return t.substring + currLongest, currHasDollar, currHasHash

}

func longestSubstring(str1, str2 string) string {
	var long, short string
	if len(str1) > len(str2) {
		long, short = str1, str2
	} else {
		long, short = str2, str1
	}

	t := buildTrie(long, nil)
	t = buildTrie(short, t)

	a, _, _ := substringFromBoth(t)

	return a

}

func buildTrie(s string, t *Trie) *Trie {
	if t == nil {
		t = &Trie{"ˆ", -1, make([]*Trie, 0)}
	}

	for i := len(s) - 1; i >= 0; i-- {
		t.insertSubstring(s[i:], i)
	}

	return t
}

func main() {
	s1 := "calavera$"
	s2 := "malavis#"

	fmt.Println(longestSubstring(s1, s2))

}
