package main

import (
	"github.com/GoncharovFyodor/fgchat/server/internal/srv"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	server := srv.NewServer()
	http.Handle("/ws", websocket.Handler(server.HandleWS))
	http.Handle("/orderbookfeed", websocket.Handler(server.HandleWSOrderbook))
	http.ListenAndServe(":3000", nil)
}
