package builder

import (
	"log"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	builder := NewConcreteBuilder()
	builderN := NewConcreteBuilderN()

	dircotor := NewDirect(&builder)
	result := dircotor.Make()
	log.Println(result.Age)

	dircotor.setBuilder(&builderN)
	resultN := dircotor.Make()
	log.Println(resultN.Age)

}
