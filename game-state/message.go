package game_state

type Message struct {
	PlayerName  string
	MessageType string

	Data map[string]any
}

func NewMessage(playerName string, messageType string, data map[string]any) Message {
	returnVal := new(Message)
	returnVal.PlayerName = playerName
	returnVal.MessageType = messageType
	returnVal.Data = data
	return *returnVal
}
