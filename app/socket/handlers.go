package socket

import (
	"backend-dinakom/configs"

	io "github.com/ambelovsky/gosf-socketio"
)

func StartConnectionHandler() {
	server := configs.IoServer
	server.On(io.OnConnection, func(c *io.Channel) {
		client := new(Client)
		client.Channel = c
	})
}
