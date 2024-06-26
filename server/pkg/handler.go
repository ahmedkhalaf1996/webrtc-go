package pkg

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoomRequestHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}

	log.Println(AllRooms.Map)
	c.JSON(http.StatusOK, resp{RoomID: roomID})
}

// JoinRoomRequestHandler handles WebSocket connections for joining a room
func JoinRoomRequestHandler(c *gin.Context) {
	roomID := c.Query("roomID")

	ws, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Web Socket Upgrade Error", err)
		return
	}

	AllRooms.InsertIntoRoom(roomID, false, ws)
	go Broadcaster(BroadcastData{RoomID: roomID, Client: ws})
}
