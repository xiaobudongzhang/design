package bridge

import "fmt"

type epson struct {

}

func (p *epson)printFile()  {
	fmt.Println("print by epson printer")
}

