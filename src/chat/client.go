package main

import (
	"github.com/gorilla/websocket"
)

type room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to other clients
	forward chan []byte
	// join is a channel for clients wishing to enter a room
	join chan *client
	// leave is a channel for clients wishing to exit the room
	leave chan *client
	// client holds all current clients in this room
	clients map[*client]bool
}

// client represents a single chatting user
type client struct {
	// socket is the web socket for this client
	socket *websocket.Conn
	// send is the channel on which messages are sent
	send chan []byte
	// room is the room that the client is chatting in
	room *room
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//joining
			r.clients[client] = true
		case client := <-r.leave:
			//leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			//forward the message to all the clients
			for client := range r.clients {
				select {
				case client.send <- msg
					//send message
				default:
					//tried to send
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
