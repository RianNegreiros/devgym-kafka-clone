package main

import (
	"bufio"
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
	server := &Server{
		Topics: make(map[string]*Topic),
	}

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, server)
	}
}

func handleConnection(conn net.Conn, server *Server) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := scanner.Text()

		switch command {
		case "PUBLISH":
			handlePublish(conn, server)
		case "CONSUME":
			handleConsume(conn, server)
		default:
			fmt.Fprintln(conn, "Unknown command:", command)
		}
	}
}

func handlePublish(conn net.Conn, server *Server) {
	fmt.Fprintln(conn, "Enter topic name:")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	topicName := scanner.Text()

	server.mu.Lock()
	topic, ok := server.Topics[topicName]
	if !ok {
		topic = &Topic{Name: topicName}
		server.Topics[topicName] = topic
	}
	server.mu.Unlock()

	fmt.Fprintln(conn, "Enter message content:")
	scanner.Scan()
	messageContent := scanner.Text()

	topic.mu.Lock()
	topic.Messages = append(topic.Messages, Message{Content: messageContent})
	topic.mu.Unlock()

	fmt.Fprintln(conn, "Message published successfully.")
}

func handleConsume(conn net.Conn, server *Server) {
	fmt.Fprintln(conn, "Enter topic name:")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	topicName := scanner.Text()

	server.mu.Lock()
	topic, ok := server.Topics[topicName]
	if !ok {
		server.mu.Unlock()
		fmt.Fprintln(conn, "Topic not found.")
		return
	}
	server.mu.Unlock()

	topic.mu.Lock()
	for _, message := range topic.Messages {
		fmt.Fprintln(conn, message.Content)
	}
	topic.mu.Unlock()

	fmt.Fprintln(conn, "Waiting for new messages...")
}
