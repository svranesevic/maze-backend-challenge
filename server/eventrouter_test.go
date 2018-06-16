package server

import (
	"sync"
	"testing"
	"time"
)

func Test_Given_RunningEventRouter_When_BroadcastEventIsReceived_Then_BroadcastEventNotificationIsEmitted(t *testing.T) {
	// Given
	eventChannel := make(chan event)
	eventNotificationChannel := make(chan eventNotification)
	shutdownChannel := make(chan struct{})

	go routeEvents(eventChannel, eventNotificationChannel, shutdownChannel)

	var wg sync.WaitGroup
	var emittedEventNotifications []eventNotification
	go func() {
		for {
			select {
			case eventNotification := <-eventNotificationChannel:
				emittedEventNotifications = append(emittedEventNotifications, eventNotification)
				wg.Done()

			case <-shutdownChannel:
				return
			}
		}
	}()

	// When
	broadcastEvent := event{SequenceID: 4242, Type: "B"}

	wg.Add(1)
	eventChannel <- broadcastEvent

	wg.Wait()
	close(shutdownChannel)

	// Then
	if len(emittedEventNotifications) != 1 {
		t.Fatalf("did not receive expected number of event notifications")
	}

	broadcastEventNotification := emittedEventNotifications[0]
	if !broadcastEventNotification.IsBroadcast {
		t.Fatalf("emitted event notification is not broadcast one")
	}

	if broadcastEventNotification.Event != broadcastEvent {
		t.Errorf("event tied to event notification is not correct. got: %v, expected: %v", broadcastEventNotification.Event, broadcastEvent)
	}
}

func Test_Given_RunningEventRouter_When_StatusEventFromUserClientWithNoFollowersIsReceived_Then_NoEventNotificationIsEmitted(t *testing.T) {
	// Given
	eventChannel := make(chan event)
	eventNotificationChannel := make(chan eventNotification)
	shutdownChannel := make(chan struct{})

	go routeEvents(eventChannel, eventNotificationChannel, shutdownChannel)

	go func() {
		for {
			select {
			case eventNotification := <-eventNotificationChannel:
				t.Fatalf("Received unexpected event notification: %v", eventNotification)

			case <-shutdownChannel:
				return
			}
		}
	}()

	// When
	eventChannel <- event{SequenceID: 12, Type: "S", FromUserID: 32}
	time.Sleep(15 * time.Millisecond)
	close(shutdownChannel)

	// Then
}
