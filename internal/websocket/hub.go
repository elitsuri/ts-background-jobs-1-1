package websocket

import "sync"

type Hub struct {
	mu      sync.RWMutex
	clients map[*Client]bool
}

func NewHub() *Hub { return &Hub{clients: make(map[*Client]bool)} }

func (h *Hub) Register(c *Client) {
	h.mu.Lock(); defer h.mu.Unlock()
	h.clients[c] = true
}

func (h *Hub) Unregister(c *Client) {
	h.mu.Lock(); defer h.mu.Unlock()
	delete(h.clients, c)
	close(c.send)
}

func (h *Hub) Broadcast(msg []byte) {
	h.mu.RLock(); defer h.mu.RUnlock()
	for c := range h.clients { select { case c.send <- msg: default: } }
}

func (h *Hub) BroadcastToUser(userID int64, msg []byte) {
	h.mu.RLock(); defer h.mu.RUnlock()
	for c := range h.clients {
		if c.UserID == userID { select { case c.send <- msg: default: } }
	}
}
