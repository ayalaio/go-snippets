package main

import (
	"strconv"
)

type SubString struct {
	line     string //[abc]
	multi    int    // 3
	children []*SubString
	res      string
}

func readOrphans(a []rune, initial int) (final int, s string) {
	i := initial
	for true {
		if i >= len(a) || a[i] < 'a' || a[i] > 'z' {
			break
		}
		s = s + string(a[i])
		i++
	}
	return i, s
}

func readParenthesis(a []rune, initial int) (final int, s string) {
	ready := 1
	i := initial + 1
	for true {
		if a[i] == '[' {
			ready++
		} else if a[i] == ']' {
			ready--
		}
		if ready > 0 {
			s = s + string(a[i])
		} else {
			break
		}
		i++
	}
	return i, s
}

func (self *SubString) Expand() {
	s := self.line
	a := []rune(s)
	var d int

	for i := 0; i < len(a); i++ {
		if a[i] >= '0' && a[i] <= '9' {
			n, _ := strconv.Atoi(string(a[i]))
			d = d*10 + n
			continue
		} else {
			if a[i] == '[' {
				posOfClosePar, newline := readParenthesis(a, i)
				nc := &SubString{newline, d, []*SubString{}, ""}
				nc.Expand()
				self.children = append(self.children, nc)
				i = posOfClosePar
				d = 0
			} else { //is just chars
				posOfNextDigit, newline := readOrphans(a, i)
				nc := &SubString{newline, 1, []*SubString{}, ""}
				self.children = append(self.children, nc)
				i = posOfNextDigit - 1
			}
		}
	}
}

func (self *SubString) DoMulti() {
	for i := 0; i < len(self.children); i++ {
		self.children[i].DoMulti()
		self.res = self.res + self.children[i].res
	}
	if len(self.children) == 0 {
		self.res = self.line
	}
	s := ""
	for i := 1; i <= self.multi; i++ {
		s = s + self.res
	}
	self.res = s
	//fmt.Println(self.res)
}

func main() {
	s := "3[cp]4[de2[rlr]]clasic"
	root := &SubString{s, 1, []*SubString{}, ""}
	root.Expand()
	root.DoMulti()
	println(root.res)
}
