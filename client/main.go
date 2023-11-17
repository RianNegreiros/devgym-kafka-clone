package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Add constants for commands and responses
const (
	PublishCommand      = "PUBLISH"
	ConsumeCommand      = "CONSUME"
	ExitCommand         = "EXIT"
	EndOfMessagesMarker = "END_OF_MESSAGES"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		printError("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	printSuccess("Connected to server.")

	for {
		fmt.Print("Enter command (PUBLISH/CONSUME/EXIT): ")
		var command string
		_, err := fmt.Scanln(&command)
		if err != nil {
			printError("Error reading user input:", err)
			continue // or return, depending on your desired behavior
		}

		switch strings.ToUpper(command) {
		case PublishCommand:
			publishMessage(conn)
		case ConsumeCommand:
			consumeMessages(conn)
		case ExitCommand:
			printSuccess("Exiting.")
			return
		default:
			printError("Unknown command:", command)
		}
	}
}

func publishMessage(conn net.Conn) {
	printPrompt("Enter topic name: ")
	var topicName string
	_, err := fmt.Scanln(&topicName)
	if err != nil {
		printError("Error reading user input:", err)
		return
	}

	printPrompt("Enter message content: ")
	var messageContent string
	_, err = fmt.Scanln(&messageContent)
	if err != nil {
		printError("Error reading user input:", err)
		return
	}

	fmt.Fprintf(conn, "PUBLISH\n")
	readServerPrompt(conn) // Read the "Enter topic name:" prompt from the server

	fmt.Fprintf(conn, "%s\n", topicName)
	readServerPrompt(conn) // Read the "Enter message content:" prompt from the server

	fmt.Fprintf(conn, "%s\n", messageContent)

	printSuccess("Server response: ", readServerPrompt(conn)) // Read the "Message published successfully." response
}

func consumeMessages(conn net.Conn) {
	printPrompt("Enter topic name: ")
	var topicName string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	topicName = scanner.Text()

	fmt.Fprintf(conn, "CONSUME\n")
	readServerPrompt(conn) // Read the "Enter topic name:" prompt from the server

	fmt.Fprintf(conn, "%s\n", topicName)

	printInfo("Messages from the server:")
	for {
		response := readServerPrompt(conn)
		if response == EndOfMessagesMarker {
			break
		}
		printSuccess(response)
	}
}

func readServerPrompt(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		printError("Error reading response:", err)
		return ""
	}
	return strings.TrimSuffix(response, "\n")
}

// Helper functions for colored output
func printPrompt(message string) {
	color.New(color.FgCyan).Printf("%s", message)
}

func printSuccess(message ...interface{}) {
	color.New(color.FgGreen).Println(message...)
}

func printInfo(message ...interface{}) {
	color.New(color.FgBlue).Println(message...)
}

func printError(message ...interface{}) {
	color.New(color.FgRed).Println(message...)
}
