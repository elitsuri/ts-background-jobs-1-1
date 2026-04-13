package websocket

import ("net/http"; "time"; "github.com/gorilla/websocket"; "log")

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Client struct {
	Hub    *Hub
	UserID int64
	conn   *websocket.Conn
	send   chan []byte
}

func NewClient(hub *Hub, w http.ResponseWriter, r *http.Request) (*Client, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil { return nil, err }
	return &Client{Hub: hub, conn: conn, send: make(chan []byte, 256)}, nil
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() { ticker.Stop(); c.conn.Close(); c.Hub.Unregister(c) }()
	for { select {
		case msg, ok := <-c.send:
			if !ok { _ = c.conn.WriteMessage(websocket.CloseMessage, []byte{}); return }
			_ = c.conn.WriteMessage(websocket.TextMessage, msg)
		case <-ticker.C: _ = c.conn.WriteMessage(websocket.PingMessage, nil)
	} }
}

func (c *Client) ReadPump() {
	defer func() { c.Hub.Unregister(c); c.conn.Close() }()
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil { log.Println("ws read:", err); break }
	}
}
