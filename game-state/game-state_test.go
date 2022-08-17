package game_state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerConnectAndDisconnect(t *testing.T) {
	gameState := NewGameState()
	message := NewMessage("JamesCameron", MessageType_Connect, map[string]any{})

	result := gameState.ProcessMessage(message)

	assert.Nil(t, result)
	assert.Equal(t, 1, len(gameState.Players))
	assert.Equal(t, 0, gameState.Players[0].id)
	assert.Equal(t, "JamesCameron", gameState.Players[0].Name)

	message = NewMessage("JackRyanJr", MessageType_Connect, map[string]any{})
	result = gameState.ProcessMessage(message)

	assert.Nil(t, result)
	assert.Equal(t, 2, len(gameState.Players))
	assert.Equal(t, 1, gameState.Players[1].id)
	assert.Equal(t, "JackRyanJr", gameState.Players[1].Name)

	result = gameState.ProcessMessage(message)
	assert.NotNil(t, result)
	assert.Equal(t, result.Error(), "a player with that name is already connected")

	message = NewMessage("", MessageType_Connect, map[string]any{})
	result = gameState.ProcessMessage(message)
	assert.NotNil(t, result)
	assert.Equal(t, "player names cannot be empty", result.Error())
	assert.Equal(t, 2, len(gameState.Players))

	message = NewMessage("NotAPlayerWhoIsConnected", MessageType_Disconnect, map[string]any{})
	result = gameState.ProcessMessage(message)
	assert.Equal(t, result.Error(), "a player who is not connected cannot disconnect")
	assert.Equal(t, 2, len(gameState.Players))

	message = NewMessage("JamesCameron", MessageType_Disconnect, map[string]any{})
	result = gameState.ProcessMessage(message)
	assert.Nil(t, result)
	assert.Equal(t, 1, len(gameState.Players))
	assert.Equal(t, "JackRyanJr", gameState.Players[0].Name)
	assert.Equal(t, 1, gameState.Players[0].id)
	// Note that the other players' ID fields remain the same when a player disconnects.
	// That player's ID will not be reassigned.
}
