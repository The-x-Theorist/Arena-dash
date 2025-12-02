package main

import (
	"log"
	"sync"
)

type PlayerInput struct {
	PlayerID string
	Seq      int
	Pressed  []string
}

type Room struct {
	ID      string
	mu      sync.Mutex
	Players map[string]*Player
	Inputs  chan PlayerInput
	Tick    int
}

func NewRoom() *Room {
	return &Room{
		ID:      GenerateRoomID(),
		Players: make(map[string]*Player),
		Inputs:  make(chan PlayerInput, 128),
	}
}

func (r *Room) AddPlayer(p *Player) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Players[p.ID] = p
	log.Printf("Player %s joined room %s", p.ID, r.ID)
}

func (r *Room) RemovePlayer(playerID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.Players, playerID)
	log.Printf("Player %s left room %s", playerID, r.ID)
}
