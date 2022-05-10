package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Store interface {
	GetPlayer(name string) (string, error)
	StorePlayer(name string) (string, error)
}

type PlayerServer struct {
	store Store
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.storePlayer(w, player)
	case http.MethodGet:
		p.showPlayer(w, player)
	}
}

func (p *PlayerServer) showPlayer(w http.ResponseWriter, player string) {
	score, _ := p.store.GetPlayer(player)

	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) storePlayer(w http.ResponseWriter, player string) {
	p.store.StorePlayer(player)

}
