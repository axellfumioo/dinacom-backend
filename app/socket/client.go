package socket

import (
	io "github.com/ambelovsky/gosf-socketio"
)

type Client struct {
	Channel *io.Channel
	Rooms   map[string]bool
}

func NewClient(c *io.Client) *Client {
	return &Client{
		Channel: &c.Channel,
		Rooms:   map[string]bool{},
	}
}
