package creational

import "testing"

func TestGetInstance(t *testing.T) {
	counter := GetInstance()
	if counter == nil {
		t.Error("expected pointer to Singleton not nil")
	}
	if counter.GetValue() != 0 {
		t.Error("Expected counter should be 0")
	}
	counter.AddOne()
	if counter.GetValue() != 1 {
		t.Error("Expected counter should be 1")
	}
	sameCounter := GetInstance()
	sameCounter.AddOne()
	if sameCounter.GetValue() != 2 {
		t.Error("Expected same counter should be 2")
	}

}
