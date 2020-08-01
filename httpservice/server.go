package main

import (
		"fmt"
		"net/http"
		"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, GetPlayerScore(player))
	
}
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
// func GetPlayerScore(name string) string {
// 	if player == "Pepper" {
// 		fmt.Fprint(w, "20")
// 		return
// 	}
// 	if player == "Floyd" {
// 		fmt.Fprint(w,"10")
// 		return
// 	}
// 	return ""
// }

func main() {
	server := &PlayerServer{}

	if err := http.ListenAndServe(":5000", server); err !=nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}