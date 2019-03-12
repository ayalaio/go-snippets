package singleton

import "sync"

type locksingleton struct {
	count int
	sync.RWMutex
}

var linstance locksingleton

func GetLInstance() *locksingleton {
	return &linstance
}

func (s *locksingleton) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *locksingleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
