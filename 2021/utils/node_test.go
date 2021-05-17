package utils

import (
	"testing"
)

func TestNode(t *testing.T) {
	n := NewNode(1)

	if n.Number != 1 {
		t.Errorf("Expecting its content to be 1")
	}

	n.AddChildren(NewNode(5))

	if len(n.Children) != 1 && n.Children[0].Number != 5 {
		t.Errorf("AddChildren is not adding children")
	}
}
