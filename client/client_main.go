package main

import (
	"fmt"

	"github.com/GoncharovFyodor/fgchat/client/internal/cl"
	"golang.org/x/net/websocket"
)

func main() {
	ws, err := websocket.Dial("ws://localhost:3000/ws", "", "http://localhost")
	if err != nil {
		fmt.Println("Failed dialing server: ", err)
	}
	c := make(chan int)

	go client.ReadLoop(ws)
	go client.WriteLoop(ws, c)

	_, ok := <-c
	if !ok {
		ws.Close()
	}
}
