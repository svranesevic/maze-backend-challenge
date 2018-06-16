package server

func dispatchEvents(eventNotificationChannel <-chan eventNotification, userClientChannel <-chan userClient, shutdownChannel <-chan struct{}) {
	userClientsEventChannel := make(map[int]chan event)

	for {
		select {
		case eventNotification := <-eventNotificationChannel:
			if eventNotification.IsBroadcast {
				for _, userClientEventChannel := range userClientsEventChannel {
					userClientEventChannel <- eventNotification.Event
				}
			} else if userClientEventChannel, ok := userClientsEventChannel[eventNotification.ToUserClientID]; ok {
				userClientEventChannel <- eventNotification.Event
			}

		case userClient := <-userClientChannel:
			eventChannel := make(chan event)
			userClientsEventChannel[userClient.ID] = eventChannel

			go func() {
				for {
					select {
					case event := <-eventChannel:
						if _, err := userClient.Connection.Write([]byte(event.RawEvent + "\r\n")); err != nil {
							log.Errorf("cold not send to client `%d`: %v\n", userClient.ID, err)
						}

					case <-shutdownChannel:
						return
					}
				}
			}()

		case <-shutdownChannel:
			return
		}
	}
}
