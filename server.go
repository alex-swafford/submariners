package main

import (
	"fmt"
	"log"
	"net/http"

	comms "submariners/server/comms"
)

func main() {
	fileServer := http.FileServer(http.Dir("./client/build"))
	http.Handle("/", fileServer)
	http.HandleFunc("/ws", comms.SocketHandler)
	fmt.Println("Serving on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
