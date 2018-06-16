package server

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

func acceptUserClients(userClientListener net.Listener, userClientChannel chan<- userClient, shutdownChannel chan struct{}) {
	for {
		connectionChannel := make(chan net.Conn)

		go func() {
			if connection, err := userClientListener.Accept(); err == nil {
				connectionChannel <- connection
			}
		}()

		select {
		case connection := <-connectionChannel:
			go handleUserClientConnection(connection, userClientChannel)

		case <-shutdownChannel:
			userClientListener.Close()
			return
		}
	}
}
func handleUserClientConnection(connection net.Conn, userClientChannel chan<- userClient) {
	b := bufio.NewReader(connection)
	userClientIDBytes, err := b.ReadString('\n')

	rawUserClientID := string(userClientIDBytes)
	rawUserClientID = strings.Trim(rawUserClientID, "\n")
	rawUserClientID = strings.Trim(rawUserClientID, "\r")

	userClientID, err := strconv.Atoi(rawUserClientID)
	if err != nil {
		return
	}

	userClient := userClient{
		ID:         userClientID,
		Connection: connection,
	}

	userClientChannel <- userClient
}
