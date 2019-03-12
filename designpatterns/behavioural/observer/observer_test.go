package observer

import (
	"fmt"
	"testing"
)

type SampleObserver struct {
	Name            string
	LastReceivedMsg string
}

func (s *SampleObserver) GetNotified(m interface{}) {
	msg, ok := m.(string)
	if ok {
		s.LastReceivedMsg = fmt.Sprintf("I %s got notified of %s", s.Name, msg)
	}
}

func TestAddObserver(t *testing.T) {
	publisher := Publisher{}
	o1 := SampleObserver{Name: "John"}
	o2 := SampleObserver{Name: "Rafa"}
	o3 := SampleObserver{Name: "Mike"}
	publisher.AddObserver(&o1)
	publisher.AddObserver(&o2)
	publisher.AddObserver(&o3)

	t.Run("List should be of size 3", func(t *testing.T) {
		if len(publisher.observerList) != 3 {
			t.Errorf("Observers were not added to list")
		}
	})

	t.Run("Observer 2 should be in the list", func(t *testing.T) {
		var found bool
		for _, o := range publisher.observerList {
			if o == &o2 {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Observers were not added to list")
		}
	})

}

func TestRemoveObserver(t *testing.T) {
	publisher := Publisher{}
	o1 := SampleObserver{Name: "John"}
	o2 := SampleObserver{Name: "Rafa"}
	o3 := SampleObserver{Name: "Mike"}
	publisher.observerList = []Observer{&o1, &o2, &o3}

	publisher.RemoveObserver(&o3)

	t.Run("List should have size of 3", func(t *testing.T) {
		if len(publisher.observerList) != 2 {
			t.Errorf("Observer was not removed from list")
		}
	})

	t.Run("o3 should not be on list", func(t *testing.T) {
		for _, o := range publisher.observerList {
			if o == &o3 {
				t.Errorf("Observer was not removed from list")
			}
		}
	})

}

func TestNotifyObservers(t *testing.T) {
	publisher := Publisher{}
	o1 := SampleObserver{Name: "John"}
	o2 := SampleObserver{Name: "Rafa"}
	o3 := SampleObserver{Name: "Mike"}
	publisher.AddObserver(&o1)
	publisher.AddObserver(&o2)
	publisher.AddObserver(&o3)
	publisher.NotifyObservers("Kawabonga!")

	t.Run("o2 Should have a message", func(t *testing.T) {
		if o2.LastReceivedMsg != "I Rafa got notified of Kawabonga!" {
			t.Errorf("Observer was not notified")
		}
	})

}
