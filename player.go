package main

import "github.com/gorilla/websocket"

type Vec2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Player struct {
	ID   string          `json:"id"`
	Name string          `json:"name"`
	Pos  Vec2            `json:"pos"`
	Vel  Vec2            `json:"-"`
	Con  *websocket.Conn `json:"-"`
}
