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
			for _, player := range gameState.Players {
				if player.name == message.playerName {
					return errors.New("a player with that name is already connected")
				}
			}
			gameState.Players = append(gameState.Players,
				NewPlayer(message.playerName, gameState.idGenerator.generateID()))
			return nil
		}
		return errors.New("player names cannot be empty")
	} else if message.messageType == MessageType_Disconnect {
		indexToRemove := -1
		for index, player := range gameState.Players {
			if player.name == message.playerName {
				indexToRemove = index
			}
		}
		if indexToRemove != -1 {
			gameState.Players = append(gameState.Players[:indexToRemove], gameState.Players[indexToRemove+1:]...)
			return nil
		} else {
			return errors.New("a player who is not connected cannot disconnect")
		}

	} else if message.messageType == MessageType_KeyPressChange {
		return errors.New("not implemented")
	} else {
		return errors.New("invalid message type")
	}
}

func NewGameState() *GameState {
	returnVal := new(GameState)
	returnVal.Players = []Player{}
	returnVal.HasStarted = false
	returnVal.idGenerator = newGuid()
	return returnVal
}
