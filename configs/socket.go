package configs

import (
	io "github.com/ambelovsky/gosf-socketio"
	transport "github.com/ambelovsky/gosf-socketio/transport"
)

var IoServer *io.Server
var IoTransport *transport.WebsocketTransport

func ConnectSocketIO() {
	if IoServer == nil && IoTransport == nil {
		// Transport configuration
		IoTransport = transport.GetDefaultWebsocketTransport()

		// SocketIo Server
		IoTransport.UnsecureTLS = true
		IoServer = io.NewServer(IoTransport)
	}
}
