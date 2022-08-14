package game_state

type Message struct {
	playerName  string
	messageType string

	data map[string]any
}

func NewMessage(playerName string, messageType string, data map[string]any) Message {
	returnVal := new(Message)
	returnVal.playerName = playerName
	returnVal.messageType = messageType
	returnVal.data = data
	return *returnVal
}
