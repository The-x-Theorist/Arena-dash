package main

type ClientMessage struct {
	Type    string   `json:"type"`
	Name    string   `json:"name,omitempty"`
	RoomID  string   `json:"roomId,omitempty"`
	Seq     int      `json:"seq,omitempty"`
	Pressed []string `json:"pressed,omitempty"`
	Height  float64  `json:"height"`
	Width   float64  `json:"width"`
}

type PlayerState struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	X             float64 `json:"x"`
	Y             float64 `json:"y"`
	OrbsCollected int     `json:"orbsCollected"`
}

type ServerWelcome struct {
	PlayerID string `json:"playerId"`
}

type ServerState struct {
	Tick    int           `json:"tick"`
	Players []PlayerState `json:"players"`
}
