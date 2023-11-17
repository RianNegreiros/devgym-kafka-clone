package main

import (
	"log"
	"net"
	"os"

	"github.com/RianNegreiros/devgym-kafka-clone/internal/client"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	c := client.NewClient(conn)
	c.Run()
}
