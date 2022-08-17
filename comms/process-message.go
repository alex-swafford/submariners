package comms

import (
	"fmt"
	game_state "submariners/server/game-state"
)

/**
 * Processes the message.
 * Returns a boolean value indicating whether to reply and the reply as a message object.
 */
func ProcessMessage(message game_state.Message, currentGameState *game_state.GameState) (sendReply bool, reply game_state.Message) {
	err := currentGameState.ProcessMessage(message)
	reply = game_state.Message{
		Data: map[string]any{},
	}
	reply.PlayerName = message.PlayerName
	if message.MessageType == game_state.MessageType_Connect {
		sendReply = true
		if err != nil {
			reply.MessageType = MessageType_ConnectionRefused
		} else {
			reply.MessageType = MessageType_Connected
			reply.Data = map[string]any{
				"connectedPlayers": gameState.Players,
			}
			fmt.Print("Players connected:")
			fmt.Println(gameState.Players)
		}
		return
	} else {
		sendReply = true
		reply.MessageType = "ERROR"
		reply.Data = map[string]any{
			"Problem": "Unknown message type: " + message.MessageType,
		}
		return
	}
}
