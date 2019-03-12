package main

import (
	"fmt"
	"time"
)

const (
	TIMER = iota + 1
	TALKER
)

type Command interface {
	Info() string
	Execute()
}

type TimerCommand struct{}

func (c *TimerCommand) Info() string {
	return fmt.Sprintf("%s", time.Now())
}

func (c *TimerCommand) Execute() {
	println("Giving Time: " + c.Info())
}

type TalkerCommand struct{}

func (c *TalkerCommand) Info() string {
	return "hello world!"
}

func (c *TalkerCommand) Execute() {
	println("Talking: " + c.Info())
}

type CommandCreator struct {
}

func (cc *CommandCreator) GetCommand(cType int) Command {
	switch cType {
	case TIMER:
		return new(TimerCommand)
	case TALKER:
		return new(TalkerCommand)
	}
	return nil
}

type CommandQueue struct {
	Queue []Command
}

func (cq *CommandQueue) AddCommand(c Command) CommandQueue {
	cq.Queue = append(cq.Queue, c)
	if len(cq.Queue) >= 3 {
		for _, command := range cq.Queue {
			command.Execute()
		}
		cq.Queue = []Command{}
	}
	return *cq
}

func main() {
	cc := CommandCreator{}
	cq := CommandQueue{}

	cq.AddCommand(cc.GetCommand(TALKER))
	cq.AddCommand(cc.GetCommand(TIMER))
	cq.AddCommand(cc.GetCommand(TALKER))
	cq.AddCommand(cc.GetCommand(TIMER))
	cq.AddCommand(cc.GetCommand(TIMER))

}
