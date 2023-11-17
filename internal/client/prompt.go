package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Client represents the client application.
type Client struct {
	conn net.Conn
}

// NewClient creates a new Client instance.
func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

// Run starts the client application.
func (c *Client) Run() {
	defer c.conn.Close()

	printSuccess("Connected to server.")

	for {
		fmt.Print("Enter command (PUBLISH/CONSUME/EXIT): ")
		var command string
		_, err := fmt.Scanln(&command)
		if err != nil {
			printError("Error reading user input:", err)
			continue
		}

		switch strings.ToUpper(command) {
		case PublishCommand:
			c.publishMessage()
		case ConsumeCommand:
			c.consumeMessages()
		case ExitCommand:
			printSuccess("Exiting.")
			return
		default:
			printError("Unknown command:", command)
		}
	}
}

func (c *Client) publishMessage() {
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

	fmt.Fprintf(c.conn, "PUBLISH\n")
	readServerPrompt(c.conn) // Read the "Enter topic name:" prompt from the server

	fmt.Fprintf(c.conn, "%s\n", topicName)
	readServerPrompt(c.conn) // Read the "Enter message content:" prompt from the server

	fmt.Fprintf(c.conn, "%s\n", messageContent)

	printSuccess("Server response: ", readServerPrompt(c.conn)) // Read the "Message published successfully." response
}

func (c *Client) consumeMessages() {
	printPrompt("Enter topic name: ")
	var topicName string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	topicName = scanner.Text()

	fmt.Fprintf(c.conn, "CONSUME\n")
	readServerPrompt(c.conn) // Read the "Enter topic name:" prompt from the server

	fmt.Fprintf(c.conn, "%s\n", topicName)

	printInfo("Messages from the server:")
	for {
		response := readServerPrompt(c.conn)
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
