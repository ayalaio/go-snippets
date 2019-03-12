package chain

import (
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (l *FirstLogger) Next(s string) {
	println(s)
	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (l *SecondLogger) Next(s string) {
	if strings.Contains(s, "hello") {
		println(s)
		if l.NextChain != nil {
			l.NextChain.Next(s)
		}
	}
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (l *WriterLogger) Next(s string) {
	if l.Writer != nil {
		l.Writer.Write([]byte(s))
	}

	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}

type ClojureLogger struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (l *ClojureLogger) Next(s string) {

	if l.Closure != nil {
		l.Closure(s)
	}

	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}
