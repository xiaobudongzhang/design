package factory

import "log"

type Restaurant interface {
	GetFood()
}

type Donglaishun struct {

}

func (* Donglaishun)GetFood()  {
	log.Println("东来顺...")
}


type Qingfeng struct {

}

func (q *Qingfeng)GetFood()  {
	log.Println("庆丰...")
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Qingfeng{}
	}
	return nil
}
