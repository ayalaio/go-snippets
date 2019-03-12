package structural

import "fmt"

type LegacyPrinter struct{}

func (lp *LegacyPrinter) Print(s string) string {
	return fmt.Sprintf("Legacy: %s", s)
}

type ModernPrinter struct {
	MessageToPrint string
}

func (mp *ModernPrinter) Print() string {
	return fmt.Sprintf("Modern: %s", mp.MessageToPrint)
}

type PrinterAdapter struct {
	LP *LegacyPrinter
	MP *ModernPrinter
}

func (pa *PrinterAdapter) Print(s string) (rt string) {
	switch {
	case pa.LP != nil:
		rt = pa.LP.Print(s)
		rt = fmt.Sprintf("Adapter: %s", rt)
	case pa.MP != nil:
		pa.MP.MessageToPrint = s
		rt = pa.MP.Print()
		rt = fmt.Sprintf("Adapter: %s", rt)
	}
	return
}
