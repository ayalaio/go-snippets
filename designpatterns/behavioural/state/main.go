package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readNumber() int {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("That is not a number, try again")
		i = readNumber()
	}
	return i
}

type State interface {
	Exec(*GameContext)
}

type GameContext struct {
	MagicNumber int
	Tries       int
}

func (g *GameContext) Run() bool {
	nextState := StartState{}
	nextState.Exec(g)
	return false
}

type StartState struct{}

func (state *StartState) Exec(g *GameContext) {
	fmt.Print("Enter number of tries: ")
	g.Tries = readNumber()

	rand.Seed(time.Now().UnixNano())
	g.MagicNumber = rand.Intn(10)
	nextState := ActiveState{}
	nextState.Exec(g)
}

type ActiveState struct{}

func (state *ActiveState) Exec(g *GameContext) {
	g.Tries = g.Tries - 1
	fmt.Println("Guess a number: ")
	guess := readNumber()
	var nextState State
	if guess == g.MagicNumber {
		nextState = &WinState{}
	} else if g.Tries > 0 {
		nextState = &ActiveState{}
	} else {
		nextState = &LoseState{}
	}
	nextState.Exec(g)
}

type WinState struct{}

func (state *WinState) Exec(g *GameContext) {
	fmt.Println("You won!")
}

type LoseState struct{}

func (state *LoseState) Exec(g *GameContext) {
	fmt.Printf("You lose! The number was %d \n", g.MagicNumber)
}

func main() {
	g := GameContext{}
	g.Run()
}
