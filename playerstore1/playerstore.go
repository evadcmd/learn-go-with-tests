package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type server struct {
	store map[string]int
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score, ok := s.store[player]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "")
	}
	fmt.Fprint(w, fmt.Sprint(score))
}

var defaultServer *server = &server{
	store: map[string]int{
		"TATA": 5,
		"RURU": 2,
	},
}

func main() {
	log.Fatal(http.ListenAndServe(":8090", defaultServer))
}
