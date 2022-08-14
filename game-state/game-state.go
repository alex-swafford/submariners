package game_state

import "errors"

type GameState struct {
	Players     []Player
	HasStarted  bool
	idGenerator *guid

	// TODO add torpedoes, treasure, etc.
}

func (gameState *GameState) processMessage(message Message) error {
	if message.messageType == MessageType_Connect {
		if message.playerName != "" {
			gameState.Players = append(gameState.Players,
				NewPlayer(message.playerName, gameState.idGenerator.generateID()))
			return nil
		}
		return errors.New("Player names cannot be empty")
	} else if message.messageType == MessageType_Disconnect {
		return errors.New("Not Implemented")
	} else if message.messageType == MessageType_KeyPressChange {
		return errors.New("Not Implemented")
	} else {
		return errors.New("Invalid message type")
	}
}

func NewGameState() *GameState {
	returnVal := new(GameState)
	returnVal.Players = []Player{}
	returnVal.HasStarted = false
	returnVal.idGenerator = newGuid()
	return returnVal
}
