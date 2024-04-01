package server

import (
	"fmt"
	"net"

	"github.com/platonoff-dev/yacache/internal/commands"
	"github.com/platonoff-dev/yacache/internal/engine"
	"github.com/platonoff-dev/yacache/internal/protocol"
)

func StartServer() {
	fmt.Printf("Starting server on port 6666\n")
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic(err)
	}

	for {
		connection, err := listener.Accept()
		fmt.Printf("Accepted connection from %s\n", connection.RemoteAddr())
		if err != nil {
			fmt.Errorf("Error accepting connection: %s", err)
			continue
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buf := make([]byte, 1024)

	for {
		n, err := connection.Read(buf)
		if err != nil {
			fmt.Errorf("Error reading from connection: %s", err)
			break
		}

		fmt.Printf("Received %d bytes: %s\n", n, buf[:n])

		message, err := protocol.Parse(buf[:n])
		if err != nil {
			connection.Write([]byte(fmt.Sprintf("-ERR %s", err)))
			break
		}

		command, err := commands.NewCommand(message)
		if err != nil {
			connection.Write([]byte(fmt.Sprintf("-ERR falied to parse command: %s", err)))
			break
		}

		result, err := engine.ExecuteCommand(command)
		if err != nil {
			connection.Write([]byte(fmt.Sprintf("-ERR falied to execute command: %s", err)))
			break
		}
		connection.Write([]byte(result))
	}

	fmt.Println("Closing connection")
}
