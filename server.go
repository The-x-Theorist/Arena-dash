package main

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type GameServer struct {
<<<<<<< HEAD
	mu    sync.Mutex
=======
	mu sync.Mutex
>>>>>>> main
	Rooms map[string]*Room
}

func NewGameServer() *GameServer {
	return &GameServer{
		Rooms: make(map[string]*Room),
	}
}

<<<<<<< HEAD
func (s *GameServer) GetOrCreateRoom(id string, height float64, width float64) *Room {
=======
func (s *GameServer) GetOrCreateRoom(id string) *Room {
>>>>>>> main
	s.mu.Lock()
	defer s.mu.Unlock()

	if r, ok := s.Rooms[id]; ok {
		return r
	}

<<<<<<< HEAD
	r := NewRoom(id, height, width)
	s.Rooms[id] = r
=======
	r := NewRoom()
	s.Rooms[r.ID] = r
>>>>>>> main

	go r.Start()
	log.Println("created room", id)
	return r
}

func (s *GameServer) Join(roomId string, name string, conn *websocket.Conn) *Player {
<<<<<<< HEAD
	s.mu.Lock()
	defer s.mu.Unlock()
=======
>>>>>>> main
	room := s.Rooms[roomId]
	playerID := GenerateRandomID()

	p := &Player{
<<<<<<< HEAD
		ID:   playerID,
=======
		ID: playerID,
>>>>>>> main
		Name: name,
		Pos: Vec2{
			X: 250,
			Y: 250,
		},
		Vel: Vec2{},
		Con: conn,
	}

	room.AddPlayer(p)
	return p
<<<<<<< HEAD
}
=======
}
>>>>>>> main
