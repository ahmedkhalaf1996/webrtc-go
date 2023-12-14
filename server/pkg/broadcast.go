package pkg

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var broadcasterStatus = make(map[string]bool)

type MessageType string

const (
	MessageTypeDefault    MessageType = "default"
	MessageTypeDisconnect MessageType = "disconnect"
)

type broadcastMsg struct {
	Message     map[string]interface{}
	RoomID      string
	Client      *websocket.Conn
	MessageType MessageType
}

var broadcast = make(chan broadcastMsg)

func BroadcastDisconnect(roomID string, sender *websocket.Conn) {
	message := broadcastMsg{
		Message:     map[string]interface{}{"disconnect": true, "userID": sender},
		RoomID:      roomID,
		MessageType: MessageTypeDisconnect,
	}

	broadcast <- message
}

type BroadcastData struct {
	RoomID string
	Client *websocket.Conn
}

// Broadcaster handles broadcasting messages to clients in a room
func Broadcaster(data BroadcastData) {
	for {
		msg := <-broadcast

		for _, client := range AllRooms.Map[data.RoomID] {
			if msg.MessageType == MessageTypeDisconnect {
				// Notify other users about the disconnection
				disconnectMessage := map[string]interface{}{
					"disconnect": true,
				}
				err := client.Conn.WriteJSON(disconnectMessage)
				if err != nil {
					log.Println("Write Error: ", err)
				}
			} else {
				// Broadcast regular messages
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Println("Write Error: ", err)
					client.Conn.Close()
				}
			}
		}
	}
}
