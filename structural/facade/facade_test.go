package facade

import "testing"

func TestCarFacade_CreateCompleteCar(t *testing.T) {
	f :=NewCarFacade()
	f.CreateCompleteCar()
}
