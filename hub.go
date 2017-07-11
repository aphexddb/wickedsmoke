package main

import (
	"bytes"
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to clients
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// NewHub creates a new instace of the websocket hub
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Run starts listening for client messages
func (h *Hub) Run() {
	log.Println("hub: running")
	for {
		select {
		case client := <-h.register:
			log.Printf("hub: client %s registered\n", client.label)
			h.clients[client] = true
		case client := <-h.unregister:
			log.Printf("hub: client %s unregistered\n", client.label)
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// Broadcast posts a message to all clients
func (h *Hub) Broadcast(message []byte) {
	message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
	h.broadcast <- message
}
