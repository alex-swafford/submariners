package game_state

import (
	"testing"
)

func TestPlayerConnect(t *testing.T) {
	gameState := NewGameState()
	message := NewMessage("JamesCameron", MessageType_Connect, map[string]any{})

	result := gameState.processMessage(message)

	assertNil(result, t)
	assertEquals(1, len(gameState.Players), t)
	assertEquals(0, gameState.Players[0].id, t)
	assertEquals("JamesCameron", gameState.Players[0].name, t)

	message = NewMessage("JackRyanJr", MessageType_Connect, map[string]any{})
	result = gameState.processMessage(message)

	assertNil(result, t)
	assertEquals(2, len(gameState.Players), t)
	assertEquals(1, gameState.Players[1].id, t)
	assertEquals("JackRyanJr", gameState.Players[1].name, t)

	message = NewMessage("", MessageType_Connect, map[string]any{})
	result = gameState.processMessage(message)
	assertNotNil(result, "result", t)
	assertEquals("player names cannot be empty", result.Error(), t)
	assertEquals(2, len(gameState.Players), t)
}
