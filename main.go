package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	server := NewGameServer()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWS(server, w, r)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println(("Arena dash server listening on port :8000"))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleWS(server *GameServer, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	log.Println("new websocket connection")

	welcome := ServerWelcome{
		PlayerID: GenerateRandomID(),
	}

	if err := conn.WriteJSON(welcome); err != nil {
		log.Println("write welcome", err)
		conn.Close()
		return
	}

	_, data, err := conn.ReadMessage()

	if err != nil {
		log.Println("read join:", err)
		conn.Close()
		return
	}

	var msg ClientMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Println("bad join message", err)
		conn.Close()
		return
	}

	if msg.Type != "join" {
		log.Println("join was not requested")
		return
	}

	roomID := msg.RoomID
	if roomID == "" {
		roomID = GenerateRandomID()
	}

	name := msg.Name
	if name == "" {
		name = "Guest"
	}

	room := server.GetOrCreateRoom(roomID, msg.Height, msg.Width)
	player := server.Join(roomID, name, conn)

	go handlePlayerMessages(room, player)
}

func handlePlayerMessages(room *Room, player *Player) {
	conn := player.Con

	defer func() {
		room.RemovePlayer(player.ID)
		conn.Close()
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("player %s disconnected %v", player.ID, err)
			return // Exit the loop when connection fails
		}

		var msg ClientMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			log.Println("bad client message", err)
			continue
		}

		switch msg.Type {
		case "input":
			room.Inputs <- PlayerInput{
				PlayerID: player.ID,
				Seq:      msg.Seq,
				Pressed:  msg.Pressed,
			}
		default:
			log.Println("unknown client message type:", msg.Type)
		}
	}
}
