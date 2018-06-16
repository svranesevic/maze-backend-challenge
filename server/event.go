package server

type event struct {
	SequenceID int
	Type       string
	FromUserID int
	ToUserID   int

	RawEvent string
}
