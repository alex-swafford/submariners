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
}
