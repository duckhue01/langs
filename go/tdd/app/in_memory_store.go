package main

import "errors"

type InMemoryStore struct {
	players []string
}

func (i *InMemoryStore) GetPlayer(name string) (string, error) {

	for _, p := range i.players {
		if p == name {
			return p, nil
		}
	}

	return "", errors.New("can not find the player")
}

func (i *InMemoryStore) StorePlayer(name string) (string, error) {

	return "", nil
}
