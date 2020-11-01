package bridge

import "fmt"

type mac struct {
	printer printer
}

func (m *mac)print()  {
	fmt.Println("print for mac")
	m.printer.printFile()
}

func (m *mac)setPrinter(p printer)  {
	m.printer = p
}
