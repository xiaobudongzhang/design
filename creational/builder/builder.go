package builder

//抽象建超者
type Builder interface {
	Build()
	GetProduct() Product
}
//指挥者
type Director struct {
	builder Builder
}
//具体建造者
type ConcreteBuilder struct {
	Age int
}
//具体建造者
type ConcreteBuilderN struct {
	Age int
}
//产品
type Product struct {
	Age int
}

func NewConcreteBuilder() ConcreteBuilder  {
	return ConcreteBuilder{}
}

func (b *ConcreteBuilder)Build()  {
	b.Age = 100
}
func (b *ConcreteBuilder) GetProduct() Product {
	return Product{b.Age}
}

func NewConcreteBuilderN() ConcreteBuilderN  {
	return ConcreteBuilderN{}
}

func (b *ConcreteBuilderN)Build()  {
	b.Age = -100
}
func (b *ConcreteBuilderN) GetProduct() Product {
	return Product{b.Age}
}

func NewDirect(b Builder) Director {
	return Director{builder:b}
}
func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director)Make()  Product {
	d.builder.Build()
	return d.builder.GetProduct()
}



