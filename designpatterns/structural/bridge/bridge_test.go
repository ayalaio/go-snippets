package structural

import (
	"errors"
	"testing"
)

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}

	err = errors.New("Content received on Writer was empty")
	return
}

func TestDellWithWriterEngine(t *testing.T) {

	testWriter := TestWriter{}

	dell := DellPrinter{
		Engine: &WriterPrintEngine{
			Writer: &testWriter,
		},
		Message: "Hello",
	}
	dell.Print()
	if testWriter.Msg != "Dell Hello" {
		t.Log(testWriter.Msg)
		t.Error("Expecting Dell Hello from printer on the Writter")
	}
}

func TestCompaqWithWriterEngine(t *testing.T) {
	testWriter := TestWriter{}

	dell := CompaqPrinter{
		Engine: &WriterPrintEngine{
			Writer: &testWriter,
		},
		Message: "Hello",
	}
	dell.Print()
	if testWriter.Msg != "Compaq Hello" {
		t.Log(testWriter.Msg)
		t.Error("Expecting Compaq Hello from printer on the Writter")
	}
}
