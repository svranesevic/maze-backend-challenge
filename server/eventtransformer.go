package server

import (
	"errors"
	"strconv"
	"strings"
)

func transformEvents(rawEventChannel <-chan string, eventChannel chan<- event, shutdownChannel chan struct{}) {
	for {
		select {
		case rawEvent := <-rawEventChannel:
			rawEvent = strings.Trim(rawEvent, "\n")
			rawEvent = strings.Trim(rawEvent, "\r")

			if event, err := unmarshalEvent(rawEvent); err == nil {
				eventChannel <- *event
			}

		case <-shutdownChannel:
			return
		}
	}
}

func unmarshalEvent(rawEvent string) (*event, error) {
	eventFragments := strings.Split(rawEvent, "|")
	if numEventFragments := len(eventFragments); numEventFragments < 2 || numEventFragments > 4 {
		log.Error("invalid event packet")
		return nil, errors.New("invalid event packet")
	}

	var event event
	var err error
	event.SequenceID, err = strconv.Atoi(eventFragments[0])
	if err != nil {
		return nil, errors.New("invalid event `sequence id`")
	}

	eventType := eventFragments[1]
	event.Type = eventType
	event.RawEvent = rawEvent

	switch eventType {
	case "F":
		fromUserId, err := strconv.Atoi(eventFragments[2])
		if err != nil {
			return nil, errors.New("invalid event `from user id`")
		}

		toUserId, err := strconv.Atoi(eventFragments[3])
		if err != nil {
			return nil, errors.New("invalid event `to user id`")
		}

		event.FromUserID = fromUserId
		event.ToUserID = toUserId
		return &event, nil

	case "U":
		fromUserId, err := strconv.Atoi(eventFragments[2])
		if err != nil {
			return nil, errors.New("invalid event `from user id`")
		}

		toUserId, err := strconv.Atoi(eventFragments[3])
		if err != nil {
			return nil, errors.New("invalid event `to user id`")
		}

		event.FromUserID = fromUserId
		event.ToUserID = toUserId
		return &event, nil

	case "B":
		return &event, nil

	case "P":
		fromUserId, err := strconv.Atoi(eventFragments[2])
		if err != nil {
			return nil, errors.New("invalid event `from user id`")
		}

		toUserId, err := strconv.Atoi(eventFragments[3])
		if err != nil {
			return nil, errors.New("invalid event `to user id`")
		}

		event.FromUserID = fromUserId
		event.ToUserID = toUserId
		return &event, nil

	case "S":
		fromUserId, err := strconv.Atoi(eventFragments[2])
		if err != nil {
			return nil, errors.New("invalid event `from user id`")
		}

		event.FromUserID = fromUserId
		return &event, nil
	}

	log.Errorf("unsupported event type `%s`\n", eventType)
	return nil, errors.New("unsupported event type `" + eventType + "`")
}
