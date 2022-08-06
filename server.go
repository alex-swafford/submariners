package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running the server.")

	fileServer := http.FileServer(http.Dir("./client/build"))

	http.Handle("/", fileServer)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
