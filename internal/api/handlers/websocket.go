package handlers

import ("net/http"; "log")
"github.com/example/ts-background-jobs-1/internal/websocket"

type WSHandler struct{ Hub *websocket.Hub }

func (h *WSHandler) Connect(w http.ResponseWriter, r *http.Request) {
	client, err := websocket.NewClient(h.Hub, w, r)
	if err != nil { log.Println("ws upgrade:", err); return }
	h.Hub.Register(client)
	go client.WritePump()
	go client.ReadPump()
}
