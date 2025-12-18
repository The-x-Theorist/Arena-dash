package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"
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
	Orb     *Orb
	Height  float64
	Width   float64
}

type TickClientResponse struct {
	Tick    int           `json:"tick"`
	Players []PlayerState `json:"players"`
	Type    string        `json:"type"`
}

func NewRoom(id string, height float64, width float64) *Room {
	return &Room{
		ID:      id,
		Players: make(map[string]*Player),
		Inputs:  make(chan PlayerInput, 128),
		Height:  height,
		Width:   width,
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

func (r *Room) Start() {
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		r.step()
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
		p.Pos.X += p.Vel.X * dt
		p.Pos.Y += p.Vel.Y * dt

		clamp(&p.Pos.X, 0, r.Width)
		clamp(&p.Pos.Y, 0, r.Height)
	}

	r.Tick++

	state := ServerState{
		Tick: r.Tick,
	}

	for _, p := range r.Players {
		state.Players = append(state.Players, PlayerState{
			ID:            p.ID,
			Name:          p.Name,
			X:             p.Pos.X,
			Y:             p.Pos.Y,
			OrbsCollected: p.OrbsCollected,
		})
	}

	r.mu.Unlock()

	r.CatchOrb()

	clientMessage := TickClientResponse{
		Tick:    state.Tick,
		Players: state.Players,
		Type:    "tick",
	}

	data, err := json.Marshal(clientMessage)

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

func (r *Room) applyInput(in PlayerInput) {
	p, ok := r.Players[in.PlayerID]

	if !ok {
		return
	}

	var v Vec2
	speed := 100.0

	for _, key := range in.Pressed {
		switch key {
		case "UP":
			v.Y -= speed
		case "DOWN":
			v.Y += speed
		case "LEFT":
			v.X -= speed
		case "RIGHT":
			v.X += speed
		}
	}
	p.Vel = v
}

func clamp(value *float64, min float64, max float64) {
	if *value < min {
		*value = min
	}

	if *value > max {
		*value = max
	}
}
