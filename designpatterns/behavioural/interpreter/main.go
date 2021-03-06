package main

import (
	"strconv"
	"strings"
)

const (
	ADD = "+"
	SUB = "-"
	MUL = "*"
	DIV = "/"
)

type Interpreter interface {
	Read() Value
}

type Value int
type MathFunc func() Value

func (v *Value) Read() Value {
	println("READING")
	return *v
}

func (f *MathFunc) Read() Value {
	println("READING")
	return (*f)()
}

func MathFuncFactory(_i, _j Interpreter, o string) MathFunc {
	switch o {
	case ADD:
		return MathFunc(func() Value {
			i := _i.Read()
			j := _j.Read()
			return Value(i + j)
		})
	case SUB:
		return MathFunc(func() Value {
			i := _i.Read()
			j := _j.Read()
			return Value(i - j)
		})
	case MUL:
		return MathFunc(func() Value {
			i := _i.Read()
			j := _j.Read()
			return Value(i * j)
		})
	case DIV:
		return MathFunc(func() Value {
			i := _i.Read()
			j := _j.Read()
			return Value(i / j)
		})
	}
	return nil
}

func isOperator(s string) bool {
	switch {
	case strings.Contains(s, ADD):
		return true
	case strings.Contains(s, SUB):
		return true
	case strings.Contains(s, MUL):
		return true
	case strings.Contains(s, DIV):
		return true
	}
	return false
}

type Stack struct {
	tokens []Interpreter
}

func (s *Stack) Push(i Interpreter) {
	s.tokens = append(s.tokens, i)
}

func (s *Stack) Pop() Interpreter {
	size := len(s.tokens)
	t := s.tokens[size-1]
	s.tokens = s.tokens[0 : size-1]
	return t
}

func strToVal(s string) Interpreter {
	i, _ := strconv.Atoi(s)
	v := Value(i)
	return &v
}

func main() {
	query := "15 7 1 1 + - / 3 * 2 1 1 + + -"
	tokens := strings.Split(query, " ")
	stack := Stack{}
	for _, token := range tokens {
		if isOperator(token) {
			right := stack.Pop()
			left := stack.Pop()
			f := MathFuncFactory(left, right, token)
			stack.Push(&f)
		} else {
			v := strToVal(token)
			stack.Push(v)
		}
	}
	println("DONE BUILDING")
	println(int(stack.Pop().Read()))
}
