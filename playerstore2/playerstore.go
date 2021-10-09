package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type httpHandler struct {
	store map[string]int
	mu    sync.RWMutex
}

func (s *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	var score int
	var ok bool
	switch r.Method {
	case http.MethodGet:
		s.mu.RLock()
		score, ok = s.store[player]
		s.mu.RUnlock()
	case http.MethodPost:
		s.mu.Lock()
		s.store[player]++
		score, ok = s.store[player]
		s.mu.Unlock()
	}
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, fmt.Sprint(score))
}

var defaultHttpHandler *httpHandler = &httpHandler{
	store: map[string]int{
		"TATA": 5,
		"RURU": 2,
	},
}

func main() {
	log.Fatal(http.ListenAndServe(":8090", defaultHttpHandler))
}
