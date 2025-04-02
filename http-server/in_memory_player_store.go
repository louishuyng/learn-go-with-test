package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.Mutex{},
	}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mutex sync.Mutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() League {
	var league []Player
	for name, score := range i.store {
		league = append(league, Player{name, score})
	}
	return league
}
