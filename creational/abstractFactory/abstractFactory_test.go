package abstractFactory

import "testing"

func TestNewSimpleLuchFacoty(t *testing.T) {
	factory := NewSimpleLuchFacoty()
	food := factory.CreateFood()
	food.Cook()

	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}