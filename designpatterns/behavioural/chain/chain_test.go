package chain

import (
	"strings"
	"testing"
)

type myTestWriter struct {
	receivedMessage *string
}

func (m *myTestWriter) Write(p []byte) (int, error) {
	tempMessage := string(p)
	m.receivedMessage = &tempMessage
	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}

func TestCreateDefaultChain(t *testing.T) {
	myWriter := myTestWriter{}
	WriterLogger := WriterLogger{Writer: &myWriter}
	second := SecondLogger{NextChain: &WriterLogger}
	chain := FirstLogger{NextChain: &second}

	t.Run("When message does not contain hello", func(t *testing.T) {

		chain.Next("message that breaks the chain")

		t.Run("Third chain wont get a variable", func(t *testing.T) {
			if myWriter.receivedMessage != nil {
				t.Error("Third chain should not receive message")
			}
		})

	})

	t.Run("When message contains hello", func(t *testing.T) {

		chain.Next("hello my friend")

		t.Run("Third chain will get a variable", func(t *testing.T) {
			if myWriter.receivedMessage == nil {
				t.Error("Third chain should receive message")
			} else if !strings.Contains(*myWriter.receivedMessage, "hello") {
				t.Error("Third chain should receive hello")
			}
		})

	})
}

func TestCreateClosureChain(t *testing.T) {
	myWriter := myTestWriter{}
	WriterLogger := WriterLogger{Writer: &myWriter}
	var receivedMsgOnClosure *string
	closureLogger := ClojureLogger{
		NextChain: &WriterLogger,
		Closure: func(s string) {
			receivedMsgOnClosure = &s
		},
	}
	chain := FirstLogger{NextChain: &closureLogger}

	t.Run("When a closure logger is in the chain", func(t *testing.T) {

		chain.Next("hello my friend")

		t.Run("Closure logger received a message", func(t *testing.T) {
			if receivedMsgOnClosure == nil {
				t.Error("Third chain should receive message")
			}
		})

	})
}
