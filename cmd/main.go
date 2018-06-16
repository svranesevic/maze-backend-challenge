package main

import (
	"os"
	"os/signal"
	"strconv"

	"maze-backend-challenge/logger"
	mazeserver "maze-backend-challenge/server"
)

var (
	log = logger.New("Main", logger.Info, os.Stdout)
)

func main() {
	configuration := obtainConfiguration()

	server, err := mazeserver.NewMazeServer(configuration)
	if err != nil {
		log.Errorf("cold not create and/or start server: %v\n", err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				server.Shutdown()
			}
		}
	}()

	select {
	case <-server.Finished:
		log.Info("server gracefully shutdown")
	}
}
func obtainConfiguration() mazeserver.Configuration {
	configuration := mazeserver.Configuration{
		EventListenerPort:      9090,
		UserClientListenerPort: 9099}

	if envEventListenerPort := os.Getenv("eventListenerPort"); envEventListenerPort != "" {
		newEventListenerPort, err := strconv.Atoi(envEventListenerPort)
		if err == nil {
			configuration.EventListenerPort = newEventListenerPort
		} else {
			log.Errorf("env variable `%s` does not represent valid port. falling back to default port: %d\n", envEventListenerPort, configuration.EventListenerPort)
		}
	}

	if envUserClientListenerPort := os.Getenv("clientListenerPort"); envUserClientListenerPort != "" {
		newUserClientListenerPort, err := strconv.Atoi(envUserClientListenerPort)
		if err == nil {
			configuration.UserClientListenerPort = newUserClientListenerPort
		} else {
			log.Errorf("env variable `%s` does not represent valid port. falling back to default port: %d\n", envUserClientListenerPort, configuration.UserClientListenerPort)
		}
	}

	return configuration
}
