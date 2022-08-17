package comms

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	game_state "submariners/server/game-state"
)

var gameState *game_state.GameState

// Must be called before processing messages.
func InitializeGameState() {
	gameState = game_state.NewGameState()
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, jsonBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Print("Received ")
		fmt.Println(string(jsonBytes))
		messageReceived := new(game_state.Message)
		err = json.Unmarshal(jsonBytes, messageReceived)
		if err != nil {
			log.Println("Failed to unmarshal JSON.")
		}

		sendReply, reply := ProcessMessage(*messageReceived, gameState)

		if sendReply {
			bytes, err := json.Marshal(reply)
			if err != nil {
				bytes = []byte("Could not marshal: " + err.Error())
			}
			if err := conn.WriteMessage(messageType, bytes); err != nil {
				log.Println(err)
				return
			}
		}

	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	reader(conn)
}
