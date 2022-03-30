package httpOperations

import (
	"bufio"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWrite(conn, message)

	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome here! %s", clientName)
	messages <- fmt.Sprintf("Client connected! %s", clientName)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- fmt.Sprintf("%s: %s", clientName, input.Text())
	}

	leavingClients <- message
	messages <- fmt.Sprintf("%s left the chat room!", clientName)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func NewConnection() {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Print(err)
	}

	go Broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}
