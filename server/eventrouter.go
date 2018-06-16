package server

type eventNotification struct {
	IsBroadcast    bool
	ToUserClientID int

	Event event
}

func routeEvents(eventChannel <-chan event, eventNotificationChannel chan<- eventNotification, shutdownChannel chan struct{}) {
	userClientFollowers := make(map[int]map[int]bool)

	for {
		select {
		case event := <-eventChannel:
			routeEvent(event, userClientFollowers, eventNotificationChannel)

		case <-shutdownChannel:
			return
		}
	}
}

func routeEvent(event event, userClientFollowers map[int]map[int]bool, eventNotificationChannel chan<- eventNotification) {
	log.Tracef("handling event: %v", event)

	switch event.Type {
	case "F":
		followers, ok := userClientFollowers[event.ToUserID]
		if !ok {
			followers = make(map[int]bool)
		}

		followers[event.FromUserID] = true
		userClientFollowers[event.ToUserID] = followers

		eventNotificationChannel <- eventNotification{IsBroadcast: false, ToUserClientID: event.ToUserID, Event: event}

	case "U":
		followers := userClientFollowers[event.ToUserID]
		delete(followers, event.FromUserID)

	case "B":
		eventNotificationChannel <- eventNotification{IsBroadcast: true, Event: event}

	case "P":
		eventNotificationChannel <- eventNotification{IsBroadcast: false, ToUserClientID: event.ToUserID, Event: event}

	case "S":
		followers := userClientFollowers[event.FromUserID]
		for follower := range followers {
			eventNotificationChannel <- eventNotification{IsBroadcast: false, ToUserClientID: follower, Event: event}
		}

	default:
		log.Warningf("unhandled event of type `%s` with sequence id `%d`\n", event.Type, event.SequenceID)
	}
}
