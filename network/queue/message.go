package queue

import "time"

// -----------------------------------------------------------------------------

// TODO
func (m *Message) TimeoutReached(d time.Duration) bool {
	if elapsed := time.Since(m.Timeout); elapsed >= d {
		return true
	}
	return false
}

// -----------------------------------------------------------------------------

// TODO
type Message struct {
	Msg     string
	Timeout time.Time
}

// NewMessage returns a new `queue.Message`.
func NewMessage(msg string) *Message { return &Message{Msg: msg, Timeout: time.Now()} }

// CONTAINER HERE
