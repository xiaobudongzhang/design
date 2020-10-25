package facade

import "log"

type CarModel struct {

}

func NewCarModel() *CarModel  {
	return &CarModel{}
}

func (c *CarModel)SetModel()  {
	log.Println("carmodel set")
}

type CarEngine struct {

}

func NewCarEngine() *CarEngine  {
	return &CarEngine{}
}

func (c *CarEngine)SetEngine()  {
	log.Println("carengin set ")
}

type CarBody struct {

}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody)SetCarBody()  {
	log.Println("carbody set")
}

type CarFacade struct {
	model *CarModel
	engine *CarEngine
	body *CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  NewCarModel(),
		engine: NewCarEngine(),
		body:   NewCarBody(),
	}
}

func (c *CarFacade)CreateCompleteCar()  {
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetCarBody()
}

