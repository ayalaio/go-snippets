package utils

type Stack []interface{}

func (self *Stack) Push(object interface{}) {
	*self = append(*self, object)
}

func (self *Stack) Pop() interface{} {
	s := *self
	l := len(s)
	o := s[l-1]
	*self = s[0 : l-1]
	return o
}
