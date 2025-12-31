package configs

import (
	io "github.com/ambelovsky/gosf-socketio"
	transport "github.com/ambelovsky/gosf-socketio/transport"
)

var ioServer *io.Server
var ioTransport *transport.WebsocketTransport

func ConnectSocketIO() {
	if ioServer == nil && ioTransport == nil {
		// Transport configuration
		ioTransport = transport.GetDefaultWebsocketTransport()

		// SocketIO Server
		ioTransport.UnsecureTLS = true
		ioServer = io.NewServer(ioTransport)
	}
}
