package server

import "testing"

func Test_Given_ValidRawFollowEvent_When_Unmarshalled_Then_ResultingEventIsValid(t *testing.T) {
	// Given
	rawFollowEvent := "666|F|60|50"

	// When
	followEvent, err := unmarshalEvent(rawFollowEvent)
	if err != nil {
		t.Fatalf("follow event could not be unmarshalled: %v", err)
	}

	// Then
	if followEvent.SequenceID != 666 {
		t.Fatalf("non matching `SequenceID`, got: %v, expected: %v", followEvent.SequenceID, 666)
	}

	if followEvent.Type != "F" {
		t.Fatalf("non matching `Type`, got: %v, expected: %v", followEvent.Type, "F")
	}

	if followEvent.FromUserID != 60 {
		t.Fatalf("non matching `FromUserID`, got: %v, expected: %v", followEvent.FromUserID, 60)
	}

	if followEvent.ToUserID != 50 {
		t.Fatalf("non matching `ToUserID`, got: %v, expected: %v", followEvent.ToUserID, 60)
	}

	if followEvent.RawEvent != rawFollowEvent {
		t.Fatalf("non matching `RawEvent`, got: %v, expected: %v", followEvent.RawEvent, rawFollowEvent)
	}
}

func Test_Given_ValidRawUnfollowEvent_When_Unmarshalled_Then_ResultingEventIsValid(t *testing.T) {
	// Given
	rawUnfollowEvent := "1|U|12|9"

	// When
	unfollowEvent, err := unmarshalEvent(rawUnfollowEvent)
	if err != nil {
		t.Fatalf("unfollow event could not be unmarshalled: %v", err)
	}

	// Then
	if unfollowEvent.SequenceID != 1 {
		t.Fatalf("non matching `SequenceID`, got: %v, expected: %v", unfollowEvent.SequenceID, 1)
	}

	if unfollowEvent.Type != "U" {
		t.Fatalf("non matching `Type`, got: %v, expected: %v", unfollowEvent.Type, "U")
	}

	if unfollowEvent.FromUserID != 12 {
		t.Fatalf("non matching `FromUserID`, got: %v, expected: %v", unfollowEvent.FromUserID, 12)
	}

	if unfollowEvent.ToUserID != 9 {
		t.Fatalf("non matching `ToUserID`, got: %v, expected: %v", unfollowEvent.ToUserID, 9)
	}
}

func Test_Given_ValidRawBroadcastEvent_When_Unmarshalled_Then_ResultingEventIsValid(t *testing.T) {
	// Given
	rawBroadcastEvent := "542532|B"

	// When
	broadcastEvent, err := unmarshalEvent(rawBroadcastEvent)
	if err != nil {
		t.Fatalf("broadcast event could not be unmarshalled: %v", err)
	}

	// Then
	if broadcastEvent.SequenceID != 542532 {
		t.Fatalf("non matching `SequenceID`, got: %v, expected: %v", broadcastEvent.SequenceID, 542532)
	}

	if broadcastEvent.Type != "B" {
		t.Fatalf("non matching `Type`, got: %v, expected: %v", broadcastEvent.Type, "B")
	}
}

func Test_Given_ValidRawPrivateMessageEvent_When_Unmarshalled_Then_ResultingEventIsValid(t *testing.T) {
	// Given
	rawPrivateMessageEvent := "43|P|32|56"

	// When
	privateMessageEvent, err := unmarshalEvent(rawPrivateMessageEvent)
	if err != nil {
		t.Fatalf("private message event could not be unmarshalled: %v", err)
	}

	// Then
	if privateMessageEvent.SequenceID != 43 {
		t.Fatalf("non matching `SequenceID`, expected got: %v, expected: %v", privateMessageEvent.SequenceID, 43)
	}

	if privateMessageEvent.Type != "P" {
		t.Fatalf("non matching `Type`, got: %v, expected: %v", privateMessageEvent.Type, "P")
	}

	if privateMessageEvent.FromUserID != 32 {
		t.Fatalf("non matching `FromUserID`, got: %v, expected: %v", privateMessageEvent.FromUserID, 32)
	}

	if privateMessageEvent.ToUserID != 56 {
		t.Fatalf("non matching `ToUserID`, got: %v, expected: %v", privateMessageEvent.ToUserID, 56)
	}
}

func Test_Given_ValidStatusUpdateEvent_When_Unmarshalled_Then_ResultingEventIsValid(t *testing.T) {
	// Given
	rawStatusUpdateEvent := "634|S|32"

	// When
	statusUpdateEvent, err := unmarshalEvent(rawStatusUpdateEvent)
	if err != nil {
		t.Fatalf("status update event could not be unmarshalled: %v", err)
	}

	// Then
	if statusUpdateEvent.SequenceID != 634 {
		t.Fatalf("non matching `SequenceID`, got: %v, expected: %v", statusUpdateEvent.SequenceID, 634)
	}

	if statusUpdateEvent.Type != "S" {
		t.Fatalf("non matching `Type`, got: %v, expected: %v", statusUpdateEvent.Type, "S")
	}

	if statusUpdateEvent.FromUserID != 32 {
		t.Fatalf("non matching `FromUserID`, got: %v, expected: %v", statusUpdateEvent.FromUserID, 32)
	}
}
