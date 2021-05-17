package utils

type Queue []interface{}

func (self *Queue) Push(o interface{}) {
	*self = append(*self, o)
}

func (self *Queue) Take() interface{} {
	l := len(*self)
	if l == 0 {
		return nil
	}
	o := (*self)[0]
	*self = (*self)[1:l]
	return o
}
