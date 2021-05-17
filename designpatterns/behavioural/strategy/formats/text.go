package formats

import (
	"fmt"

	"github.com/daroay/behavioural/strategy"
)

type TextSquare struct {
	strategy.PrintOutput
}

func (t *TextSquare) Print() error {
	if t.Writer == nil {
		return fmt.Errorf("No Writer Implemented")
	}
	t.Writer.Write([]byte("Square\n"))
	if t.LogWriter != nil {
		t.LogWriter.Write([]byte("Data was written\n"))
	}
	return nil
}
