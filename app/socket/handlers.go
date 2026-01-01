package socket

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/configs"
	"fmt"
	"strings"

	io "github.com/ambelovsky/gosf-socketio"
)

func StartConnectionHandler() {
	server := configs.IoServer

	server.On(io.OnConnection, func(c *io.Channel) {
		userID := getUserID(c)

		c.Join("user:" + userID)
		fmt.Println("Client Connected")
	})

	server.On(io.OnDisconnection, func(c *io.Channel) {
		fmt.Println("Client Disconnected")
	})

	server.On("join:room", func(c *io.Channel, roomID string) {
		c.Join("room:" + roomID)
		fmt.Println("Success to join room:" + roomID)
	})
}

func getUserID(c *io.Channel) string {
	token := c.Request().Header.Get("Authorization")
	parts := strings.Split(token, " ")
	if len(parts) != 2 {
		fmt.Println("Invalid Authorization header format")
		return ""
	}

	token = parts[1]

	fmt.Println("Token received")

	claims, err := helpers.ValidateToken(token)
	if err != nil {
		fmt.Println("Token invalid:", err.Error())
		c.Close()
		return ""
	}

	return claims.UserID
}
