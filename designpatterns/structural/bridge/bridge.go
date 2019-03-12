package structural

import (
	"fmt"
	"io"
)

type PrintEngine interface {
	PrintMessage(string)
}

type ConsolePrintEngine struct {
}

func (c *ConsolePrintEngine) PrintMessage(s string) {
	fmt.Println(s)
}

type WriterPrintEngine struct {
	Writer io.Writer
}

func (c *WriterPrintEngine) PrintMessage(s string) {
	fmt.Fprint(c.Writer, s)
}

// You can see that the actual print capability is decoupled
// from the printer and implemented on an Engine, which will be pluged
// so we can change what a printer Does as much as we want

type Printer interface {
	Print()
}

type DellPrinter struct {
	Message string
	Engine  PrintEngine
}

func (p *DellPrinter) Print() {
	p.Engine.PrintMessage(fmt.Sprintf("Dell %s", p.Message))
}

type CompaqPrinter struct {
	Message string
	Engine  PrintEngine
}

func (p *CompaqPrinter) Print() {
	p.Engine.PrintMessage(fmt.Sprintf("Compaq %s", p.Message))
}
