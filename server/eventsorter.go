package server

func sortEvents(eventChannel <-chan event, sortedEventChannel chan<- event, shutdownChannel <-chan struct{}) {
	nextEventSequenceID := 1
	eventQueue := make(map[int]event)

	for {
		select {
		case event := <-eventChannel:
			eventQueue[event.SequenceID] = event
			for {
				event, ok := eventQueue[nextEventSequenceID]
				if !ok {
					break
				}

				delete(eventQueue, event.SequenceID)
				sortedEventChannel <- event
				nextEventSequenceID++
			}

		case <-shutdownChannel:
			return
		}
	}
}
