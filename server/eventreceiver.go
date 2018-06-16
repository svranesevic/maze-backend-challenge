package server

import (
	"bufio"
	"io"
	"net"
)

func receiveEvents(eventListener net.Listener, eventChannel chan<- string, shutdownChannel chan struct{}) {
	connectionChannel := make(chan net.Conn)

	go func() {
		if connection, err := eventListener.Accept(); err == nil {
			connectionChannel <- connection
		}
	}()

	select {
	case connection := <-connectionChannel:
		b := bufio.NewReader(connection)
		for {
			rawEvent, err := b.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}

			eventChannel <- rawEvent
		}

	case <-shutdownChannel:
		eventListener.Close()
		return
	}
}
