package game_state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerConnect(t *testing.T) {
	gameState := NewGameState()
	message := NewMessage("JamesCameron", MessageType_Connect, map[string]any{})

	result := gameState.processMessage(message)

	assert.Nil(t, result)
	assert.Equal(t, 1, len(gameState.Players))
	assert.Equal(t, 0, gameState.Players[0].id)
	assert.Equal(t, "JamesCameron", gameState.Players[0].name)

	message = NewMessage("JackRyanJr", MessageType_Connect, map[string]any{})
	result = gameState.processMessage(message)

	assert.Nil(t, result)
	assert.Equal(t, 2, len(gameState.Players))
	assert.Equal(t, 1, gameState.Players[1].id)
	assert.Equal(t, "JackRyanJr", gameState.Players[1].name)

	message = NewMessage("", MessageType_Connect, map[string]any{})
	result = gameState.processMessage(message)
	assert.NotNil(t, result)
	assert.Equal(t, "player names cannot be empty", result.Error())
	assert.Equal(t, 2, len(gameState.Players))
}
