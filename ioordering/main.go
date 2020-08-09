package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	if err := http.ListenAndServe(":5000", server); err !=nil {
		log.Fatalf("couldn't listen on port 5000")
	}
}