package client

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/websocket"
)

func ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		fmt.Print("> ")

		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("Error while reading: ", err)
			continue
		}

		msg := buf[:n]
		fmt.Printf("\r%s\n", msg)
	}
}

func WriteLoop(ws *websocket.Conn, c chan<- int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("> ")
		s := scanner.Text()

		// Exit command
		if s == "q" {
			c <- 1
			close(c)
			break
		}

		ws.Write([]byte(s))
	}
}
