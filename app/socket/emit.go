package socket

import "backend-dinakom/configs"

func EmitToUser(UserID string, event string, data interface{}) {
	configs.IoServer.BroadcastTo("user:"+UserID, event, data)
}

func EmitToRoom(RoomID string, event string, data interface{}) {
	configs.IoServer.BroadcastTo("room:"+RoomID, event, data)
}
