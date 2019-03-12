package observer

type Observer interface {
	GetNotified(interface{})
}

type Publisher struct {
	observerList []Observer
}

func (s *Publisher) AddObserver(o Observer) {
	s.observerList = append(s.observerList, o)
}

func (s *Publisher) RemoveObserver(o Observer) {
	var idx int
	for i, observing := range s.observerList {
		if o == observing {
			idx = i
			break
		}
	}
	s.observerList = append(s.observerList[:idx], s.observerList[idx+1:]...)
}

func (s *Publisher) NotifyObservers(m interface{}) {
	for _, observing := range s.observerList {
		observing.GetNotified(m)
	}
}
