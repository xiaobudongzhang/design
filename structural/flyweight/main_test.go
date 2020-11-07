package flyweight

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	newPlayer("t",TerroristDressType)
	newPlayer("ct",CounterTerroristDressType)
	newPlayer("t",TerroristDressType)
	newPlayer("ct",CounterTerroristDressType)
	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
