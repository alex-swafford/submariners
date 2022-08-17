package game_state

type Player struct {
	Name string `json:"name"`
	id   int

	// TODO add position, speed, etc
}

func NewPlayer(name string, id int) Player {
	returnVal := new(Player)
	returnVal.id = id
	returnVal.Name = name

	return *returnVal
}
