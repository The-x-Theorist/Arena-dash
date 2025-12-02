package main

import (
<<<<<<< Updated upstream
	"log"
	"sync"
=======
	"encoding/json"
	"log"
	"sync"
	"time"
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
=======

func (r *Room) Start() {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		// r.step()
	}
}

func (r *Room) step() {
	r.mu.Lock()

	for {
		select {
		case in := <-r.Inputs: 
			r.applyInput(in)
		default: 
			goto doneInputs
		}
	}

	doneInputs:

	dt := 0.05
	for _, p := range r.Players {
		p.Pos.X += p.Vel.X + dt
		p.Pos.Y += p.Vel.Y + dt

		clamp(&p.Pos.X, 0, 500)
		clamp(&p.Pos.Y, 0, 500)
	}

	r.Tick++

	state := ServerState{
		Tick: r.Tick,
	}

	for _, p := range r.Players {
		state.Players = append(state.Players, PlayerState{
			ID: p.ID,
			Name: p.Name,
			X: p.Pos.X,
			Y: p.Pos.Y,
		})
	}

	r.mu.Unlock()

	data, err := json.Marshal(state)

	if err != nil {
		log.Println("Marshal state", err)
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for id, p := range r.Players {
		if p.Con == nil {
			continue
		}

		if err := p.Con.WriteMessage(1, data); err != nil {
			log.Printf("Write to player %s failed %v (removing)", id, err)
			p.Con.Close()
			delete(r.Players, id)
		}
	}
}

func clamp(value *float64, min float64, max float64) {
	if (*value < min) {
		*value = min
	}

	if (*value > max) {
		* value = max
	}
}
>>>>>>> Stashed changes
