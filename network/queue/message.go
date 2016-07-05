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
	ID      int
	Msg     string
	Timeout time.Time
}

// NewMessage returns a new `queue.Message`.
func NewMessage(ID int, msg string) *Message { return &Message{ID: ID, Msg: msg, Timeout: time.Now()} }
