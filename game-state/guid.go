package game_state

/**
* Used to create unique identifiers.
* Simple incrementing counter.
 */
type guid struct {
	nextID int
}

func (g *guid) generateID() int {
	id := g.nextID
	g.nextID++
	return id
}

func newGuid() *guid {
	returnVal := new(guid)
	returnVal.nextID = 0
	return returnVal
}
