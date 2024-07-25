package writer

import "github.com/OlyMahmudMugdho/go-dependency-injection/printer"

type Writer struct {
	PrinterMachine printer.Printer
}

func NewWriter(PrinterMachine printer.Printer) *Writer {
	return &Writer{
		PrinterMachine: PrinterMachine,
	}
}

func (w *Writer) Write() string {
	return "writing in " + w.PrinterMachine.Print() + " color"
}
