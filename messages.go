package main

type ClientMessage struct {
	Type    string   `json:"type"`
	Name    string   `json:"name,omitempty"`
	RoomID  string   `json:"roomId,omitempty"`
	Seq     int      `json:"seq,omitempty"`
	Pressed []string `json:"pressed,omitempty"`
}

type PlayerState struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

type ServerWelcome struct {
	PlayerID string `json:"playerId"`
	RoomID   string `json:"roomId"`
}

type ServerState struct {
	Tick    int           `json:"tick"`
	Players []PlayerState `json:"players"`
}
