package flyweight

type player struct {
	dress dress
	playerType string
	lat int
	long int
}

func newPlayer(playType, dressType string)*player  {
	dress,_ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType:playType,
		dress:    dress,
	}
}

func (p *player)newLocation(lat, long int)  {
	p.lat = lat
	p.long = long
}
