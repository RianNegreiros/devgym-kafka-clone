package main

import (
	"fmt"
	"net"
	"sync"
)

// Message struct represents a message in a topic
type Message struct {
	Content string
}

// Topic struct represents a topic containing messages
type Topic struct {
	Name     string
	Messages []Message
	mu       sync.Mutex
}

// Server struct represents the TCP server
type Server struct {
	Topics map[string]*Topic
	mu     sync.Mutex
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on localhost:8080")
}
