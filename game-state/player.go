package game_state

type Player struct {
	name string
	id   int

	// TODO add position, speed, etc
}

func NewPlayer(name string, id int) Player {
	returnVal := new(Player)
	returnVal.id = id
	returnVal.name = name

	return *returnVal
}
