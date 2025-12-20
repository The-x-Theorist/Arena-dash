package main

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type GameServer struct {
	mu    sync.Mutex
	Rooms map[string]*Room
}

func NewGameServer() *GameServer {
	return &GameServer{
		Rooms: make(map[string]*Room),
	}
}

func (s *GameServer) GetOrCreateRoom(id string, height float64, width float64) *Room {
	s.mu.Lock()
	defer s.mu.Unlock()

	if r, ok := s.Rooms[id]; ok {
		return r
	}

	r := NewRoom(id, height, width)
	r.SpawnNewOrb()

	s.Rooms[id] = r

	go r.Start()
	log.Println("created room", id)
	return r
}

func (s *GameServer) Join(roomId string, name string, conn *websocket.Conn) *Player {
	s.mu.Lock()
	defer s.mu.Unlock()
	room := s.Rooms[roomId]
	playerID := GenerateRandomID()

	PlayerXPos := room.Width / 2
	PlayerYPos := room.Height * 0.75

	if len(s.Rooms[roomId].Players) > 0 {
		PlayerYPos = room.Height * 0.25
	}

	p := &Player{
		ID:   playerID,
		Name: name,
		Pos: Vec2{
			X: PlayerXPos,
			Y: PlayerYPos,
		},
		Vel: Vec2{},
		Con: conn,
	}

	room.AddPlayer(p)
	return p
}
