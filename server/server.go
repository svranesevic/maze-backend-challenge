package server

import (
	"maze-backend-challenge/logger"
	"net"
	"os"
	"strconv"
)

var (
	log = logger.New("Server", logger.Info, os.Stdout)
)

type Configuration struct {
	EventListenerPort      int
	UserClientListenerPort int
}

type Server struct {
	Finished chan struct{}

	isRunning bool

	userListener  net.Listener
	eventListener net.Listener
}

type userClient struct {
	ID         int
	Connection net.Conn
}

// Create, and start new server
func NewMazeServer(configuration Configuration) (*Server, error) {
	log.Info("starting server")

	eventListener, err := net.Listen("tcp", ":"+strconv.Itoa(configuration.EventListenerPort))
	if err != nil {
		return nil, err
	}

	userClientListener, err := net.Listen("tcp", ":"+strconv.Itoa(configuration.UserClientListenerPort))
	if err != nil {
		return nil, err
	}

	shutdownChannel := make(chan struct{})

	// Accepts user client connections
	userClientChannel := make(chan userClient)
	go acceptUserClients(userClientListener, userClientChannel, shutdownChannel)

	// Receives raw events
	rawEventChannel := make(chan string)
	go receiveEvents(eventListener, rawEventChannel, shutdownChannel)

	// Transforms raw events into typed events
	eventChannel := make(chan event)
	go transformEvents(rawEventChannel, eventChannel, shutdownChannel)

	// Ensures order of events
	sortedEventChannel := make(chan event)
	go sortEvents(eventChannel, sortedEventChannel, shutdownChannel)

	// Route events
	eventNotificationChannel := make(chan eventNotification)
	go routeEvents(sortedEventChannel, eventNotificationChannel, shutdownChannel)

	// Dispatch event(s) to appropriate client(s)
	go dispatchEvents(eventNotificationChannel, userClientChannel, shutdownChannel)

	return &Server{Finished: shutdownChannel, isRunning: true, userListener: userClientListener, eventListener: eventListener}, nil
}

// Gracefully shutdown server
func (s *Server) Shutdown() error {
	if !s.isRunning {
		return nil
	}

	s.isRunning = false

	close(s.Finished)
	s.userListener.Close()
	s.eventListener.Close()
	return nil
}
