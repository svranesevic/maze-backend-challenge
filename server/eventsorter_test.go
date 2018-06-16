package server

import (
	"sync"
	"testing"
)

func Test_Given_RunningEventSorter_When_EventsAreReceivedOutOfOrder_Then_EventsAreForwarderInOrder(t *testing.T) {
	// Given
	eventChannel := make(chan event)
	sortedEventChannel := make(chan event)
	shutdownChannel := make(chan struct{})

	go sortEvents(eventChannel, sortedEventChannel, shutdownChannel)

	var wg sync.WaitGroup
	var sortedEvents []event
	go func() {
		for {
			select {
			case event := <-sortedEventChannel:
				sortedEvents = append(sortedEvents, event)
				wg.Done()

			case <-shutdownChannel:
				return
			}
		}
	}()

	// When
	wg.Add(3)
	eventChannel <- event{SequenceID: 2, Type: "B", RawEvent: "2|B"}
	eventChannel <- event{SequenceID: 1, Type: "U", FromUserID: 12, ToUserID: 9, RawEvent: "1|U|12|9"}
	eventChannel <- event{SequenceID: 3, Type: "P", FromUserID: 32, ToUserID: 56, RawEvent: "3|P|32|56"}

	wg.Wait()
	close(shutdownChannel)

	// Then
	if len(sortedEvents) != 3 {
		t.Errorf("did not receive expected number of events")
		return
	}

	nextEventSequenceID := 1
	for i := 0; i < 3; i++ {
		event := sortedEvents[i]
		if event.SequenceID != nextEventSequenceID {
			t.Errorf("received event out of order. got sequence id: %v, expected sequence id: %v", event.SequenceID, nextEventSequenceID)
			return
		}
		nextEventSequenceID++
	}
}

func Test_Given_RunningEventSorter_When_EventsAreReceivedInOrder_Then_EventsAreForwarderInOrder(t *testing.T) {
	// Given
	eventChannel := make(chan event)
	sortedEventChannel := make(chan event)
	shutdownChannel := make(chan struct{})

	go sortEvents(eventChannel, sortedEventChannel, shutdownChannel)

	var wg sync.WaitGroup

	var sortedEvents []event
	go func() {
		for {
			select {
			case event := <-sortedEventChannel:
				sortedEvents = append(sortedEvents, event)
				wg.Done()

			case <-shutdownChannel:
				return
			}
		}
	}()

	// When
	wg.Add(3)
	eventChannel <- event{SequenceID: 1, Type: "U", FromUserID: 12, ToUserID: 9, RawEvent: "1|U|12|9"}
	eventChannel <- event{SequenceID: 2, Type: "B", RawEvent: "2|B"}
	eventChannel <- event{SequenceID: 3, Type: "P", FromUserID: 32, ToUserID: 56, RawEvent: "3|P|32|56"}

	wg.Wait()
	close(shutdownChannel)

	// Then
	if len(sortedEvents) != 3 {
		t.Fatalf("did not receive expected number of events")
	}

	nextEventSequenceID := 1
	for i := 0; i < 3; i++ {
		event := sortedEvents[i]
		if event.SequenceID != nextEventSequenceID {
			t.Fatalf("received event out of order. got sequence id: %v, expected sequence id: %v", event.SequenceID, nextEventSequenceID)
		}
		nextEventSequenceID++
	}
}
