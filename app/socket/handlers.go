package socket

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/configs"
	"fmt"
	"log"
	"net/http"

	io "github.com/ambelovsky/gosf-socketio"
)

func StartConnectionHandler() {
	server := configs.IoServer

	server.On(io.OnConnection, func(c *io.Channel) {
		userID := getUserID(c)
		if userID == "" {
			return
		}

		c.Join("user:" + userID)
		fmt.Println("Client Connected")
	})

	server.On(io.OnDisconnection, func(c *io.Channel) {
		fmt.Println("Client Disconnected")
	})

	server.On("join:room", func(c *io.Channel, roomID string, ack func()) {
		fmt.Println("Joined room: room:" + roomID)
		c.Join("room:" + roomID)
		ack()
	})

	server.On("leave:room", func(c *io.Channel, roomId string) {
		fmt.Println("Leaved room: room:" + roomId)
		c.Leave("room:" + roomId)
	})

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", server)

	log.Println("Socket.IO running on :8001")
	go http.ListenAndServe(":8001", mux)
}

func getUserID(c *io.Channel) string {
	params := c.Request().URL.Query()
	token := params.Get("token")

	fmt.Println("Token received")

	claims, err := helpers.ValidateToken(token)
	if err != nil {
		fmt.Println("Token invalid:", err.Error())
		return ""
	}

	return claims.UserID
}
