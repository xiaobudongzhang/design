package abstractFactory

import "log"

type Lunch interface {
	Cook()
}

type Rise struct {

}

func (r *Rise)Cook()  {
	log.Println("rise ...")
}

type Tomato struct {

}

func (t *Tomato)Cook()  {
	log.Println("tomato")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct {
	
}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}

func NewSimpleLuchFacoty()  LunchFactory {
	return &SimpleLunchFactory{}
}

