package server

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/platonoff-dev/yacache/internal/engine"
	"github.com/platonoff-dev/yacache/internal/engine/commands"
	"github.com/platonoff-dev/yacache/internal/protocol"
)

const (
	BufferSize = 1024
)

func StartServer() {
	fmt.Printf("Starting server on port 6666\n")
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(0)
	}()

	for {
		connection, err := listener.Accept()
		fmt.Printf("Accepted connection from %s\n", connection.RemoteAddr())
		if err != nil {
			fmt.Printf("Error accepting connection: %s", err)
			continue
		}

		go handleConnection(connection)
	}
}

func response(connection net.Conn, data any) {
	message, err := protocol.Serialize(data)
	if err != nil {
		fmt.Printf("failed to build response message: %s\n", err)
		return
	}

	_, err = connection.Write(message)
	if err != nil {
		fmt.Printf("failed to write response: %s\n", err)
	}
}

func parseCommand(message any) (commands.Command, error) {
	var command []string

	rawCommand, ok := message.([]any)
	if !ok {
		return nil, errors.New("message is not and array")
	}

	for _, rawPart := range rawCommand {
		part, ok := rawPart.(string)
		if !ok {
			return nil, fmt.Errorf("not a string: %v", rawPart)
		}

		command = append(command, part)
	}

	return command, nil
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buf := make([]byte, BufferSize)
	engine := engine.New()

	for {
		n, err := connection.Read(buf)
		if err != nil {
			fmt.Printf("Error reading from connection: %s", err)
			break
		}

		fmt.Printf("Received %d bytes: %s\n", n, buf[:n])

		message, err := protocol.Parse(buf[:n])
		if err != nil {
			response(connection, fmt.Errorf("failed to parse message: %w", err))
			continue
		}

		command, err := parseCommand(message)
		if err != nil {
			response(connection, fmt.Errorf("failed to parse command: %s", err))
			continue
		}

		result, err := engine.ExecuteCommand(command)
		if err != nil {
			response(connection, fmt.Errorf("failed to execute commad: %w", err))
			continue
		}

		response(connection, result)
	}

	fmt.Println("Closing connection")
}
