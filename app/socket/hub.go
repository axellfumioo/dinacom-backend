package socket

type Hub struct {
	Rooms map[string]map[*Client]bool
}

var HubInstance = &Hub{
	Rooms: map[string]map[*Client]bool{},
}

func (h *Hub) Join(client *Client, room string) {
	if h.Rooms[room] == nil {
		h.Rooms[room] = make(map[*Client]bool)
	}
	h.Rooms[room][client] = true
	client.Rooms[room] = true
}

func (h *Hub) Broadcast(room string, event string, msg interface{}) {
	for c := range h.Rooms[room] {
		c.Channel.Emit(event, msg)
	}
}
