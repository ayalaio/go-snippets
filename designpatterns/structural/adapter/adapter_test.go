package structural

import (
	"strings"
	"testing"
)

func TestAdapter(t *testing.T) {
	legacyPrinter := LegacyPrinter{}

	adapter := PrinterAdapter{LP: &legacyPrinter}
	msg := adapter.Print("Hello")
	t.Log(msg)
	if !strings.Contains(msg, "Adapter: Legacy: Hello") {
		t.Error("Adapter not working for legacy printer")
	}

	modernPrinter := ModernPrinter{}
	adapter = PrinterAdapter{MP: &modernPrinter}
	msg = adapter.Print("Hello")
	t.Log(msg)
	if !strings.Contains(msg, "Adapter: Modern: Hello") {
		t.Error("Adapter not working for modern printer")
	}
}
