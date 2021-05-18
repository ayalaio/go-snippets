package utils

type Stack []interface{}

func NewStack() *Stack {
	return &Stack{}
}

func (self *Stack) Push(object interface{}) {
	*self = append(*self, object)
}

func (self *Stack) Pop() interface{} {
	l := len(*self)
	if l == 0 {
		return nil
	}
	o := (*self)[l-1]
	*self = (*self)[0 : l-1]
	return o
}
