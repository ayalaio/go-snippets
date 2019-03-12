package creational

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func (s *singleton) AddOne() {
	s.count += 1
}

func (s *singleton) GetValue() int {
	return s.count
}
