package main

import (
	"flag"
	"log"
	"os"

	"github.com/daroay/behavioural/strategy/formats"
)

var output = flag.String("output", "text", "...")

func main() {
	flag.Parse()
	activeStrategy, err := formats.NewPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}
	activeStrategy.SetLog(os.Stdout)
	switch *output {
	case formats.TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case formats.IMAGE_STRATEGY:
		w, err := os.Create("/tmp/shape.jpeg")
		if err != nil {
			log.Fatal("Error opening img")
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}

	if err = activeStrategy.Print(); err != nil {
		log.Fatal(err)
	}
}
