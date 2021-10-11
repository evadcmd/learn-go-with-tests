package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"sync"
)

var defaultPlayerStore = map[string]int{
	"TATA": 5, "RURU": 2,
}

type playerServeMux struct {
	store map[string]int
	mu    sync.RWMutex
	http.ServeMux
}

func (m *playerServeMux) scoreHandler(w http.ResponseWriter, r *http.Request) {
	player := path.Base(r.URL.Path)
	var score int
	var ok bool
	switch r.Method {
	case http.MethodGet:
		m.mu.RLock()
		score, ok = m.store[player]
		m.mu.RUnlock()
	case http.MethodPost:
		m.mu.Lock()
		m.store[player]++
		score, ok = m.store[player]
		m.mu.Unlock()
	}
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, fmt.Sprint(score))
}

func (m *playerServeMux) leagueHandler(w http.ResponseWriter, r *http.Request) {
	jsonb, err := json.Marshal(m.store)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprint(w, string(jsonb))
}

func NewPlayerServeMux() *playerServeMux {
	serveMux := &playerServeMux{store: defaultPlayerStore}
	serveMux.HandleFunc("/league", serveMux.leagueHandler)
	serveMux.HandleFunc("/players/", serveMux.scoreHandler)
	return serveMux
}

func main() {
	log.Fatal(http.ListenAndServe(":8090", NewPlayerServeMux()))
}
