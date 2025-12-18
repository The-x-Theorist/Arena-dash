package main

import (
	"math"
	"math/rand"
)

type Orb struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (r *Room) SpawnNewOrb() {
	r.Orb = &Orb{
		X: rand.Float64() * 500,
		Y: rand.Float64() * 500,
	}
}

func (r *Room) CatchOrb() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Orb == nil {
		return
	}

	collisionRadius := 20.0 // Adjust this value based on your orb/player size

	for _, player := range r.Players {
		distance := calculateDistance(player.Pos.X, player.Pos.Y, r.Orb.X, r.Orb.Y)

		if distance < collisionRadius {
			// Collision detected! Spawn a new orb
			r.SpawnNewOrb()
			player.OrbsCollected++
			return // Only one player can catch the orb at a time
		}
	}
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}
